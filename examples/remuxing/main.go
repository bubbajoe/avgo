package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/bubbajoe/avgo"
)

var (
	input  = flag.String("i", "", "the input path")
	output = flag.String("o", "", "the output path")
)

func main() {
	// Handle ffmpeg logs
	avgo.SetLogLevel(avgo.LogLevelInfo)
	avgo.SetLogCallback(func(l avgo.LogLevel, fmt, msg, parent string) {
		log.Printf("ffmpeg log: %s (level: %s)\n", strings.TrimSpace(msg), l.String())
	})

	// Parse flags
	flag.Parse()

	// Usage
	if *input == "" || *output == "" {
		log.Println("Usage: <binary path> -i <input path> -o <output path>")
		return
	}

	// Alloc packet
	pkt := avgo.AllocPacket()
	defer pkt.Free()

	// Alloc input format context
	inputFormatContext := avgo.AllocFormatContext()
	if inputFormatContext == nil {
		log.Fatal(errors.New("main: input format context is nil"))
	}
	defer inputFormatContext.Free()

	dict := avgo.NewDictionary()
	defer dict.Free()
	dict.Set("safe", "0", avgo.DictionaryFlags(avgo.DictionaryFlagAppend))
	// dict.Set("use_wallclock_as_timestamps", "1", avgo.DictionaryFlags(avgo.DictionaryFlagAppend))

	// Open input
	if err := inputFormatContext.OpenInput(*input, nil, dict); err != nil {
		log.Fatal(fmt.Errorf("main: opening input failed: %w", err))
	}
	defer inputFormatContext.CloseInput()

	// Find stream info
	if err := inputFormatContext.FindStreamInfo(nil); err != nil {
		log.Fatal(fmt.Errorf("main: finding stream info failed: %w", err))
	}
	format := ""
	// Alloc output format context
	if strings.HasPrefix(*output, "rtmp://") {
		format = "flv"
	}
	outputFormatContext, err := avgo.AllocOutputFormatContext(nil, format, *output)
	if err != nil {
		log.Fatal(fmt.Errorf("main: allocating output format context failed: %w", err))
	}
	if outputFormatContext == nil {
		log.Fatal(errors.New("main: output format context is nil"))
	}
	defer outputFormatContext.Free()

	// Loop through streams
	inputStreams := make(map[int]*avgo.Stream)  // Indexed by input stream index
	outputStreams := make(map[int]*avgo.Stream) // Indexed by input stream index
	for _, is := range inputFormatContext.Streams() {
		// Only process audio or video
		if is.CodecParameters().MediaType() != avgo.MediaTypeAudio &&
			is.CodecParameters().MediaType() != avgo.MediaTypeVideo {
			continue
		}

		// Add input stream
		inputStreams[is.Index()] = is

		// Add stream to output format context
		os := outputFormatContext.NewStream(nil)
		if os == nil {
			log.Fatal(errors.New("main: output stream is nil"))
		}

		// Copy codec parameters
		if err = is.CodecParameters().Copy(os.CodecParameters()); err != nil {
			log.Fatal(fmt.Errorf("main: copying codec parameters failed: %w", err))
		}

		// Reset codec tag
		os.CodecParameters().SetCodecTag(0)

		// Add output stream
		outputStreams[is.Index()] = os
	}

	// If this is a file, we need to use an io context
	if !outputFormatContext.OutputFormat().Flags().Has(avgo.IOFormatFlagNofile) {
		// Create io context
		ioContext := avgo.NewIOContext()

		// Open io context
		if err = ioContext.Open(*output, avgo.NewIOContextFlags(avgo.IOContextFlagWrite)); err != nil {
			log.Fatal(fmt.Errorf("main: opening io context failed: %w", err))
		}
		defer ioContext.Closep() //nolint:errcheck

		// Update output format context
		outputFormatContext.SetPb(ioContext)
	}

	// Write header
	if err = outputFormatContext.WriteHeader(nil); err != nil {
		log.Fatal(fmt.Errorf("main: writing header failed: %w", err))
	}

	// sps := 1024 * 1024
	// css := 0

	type timestampRef struct {
		dts      int64
		pts      int64
		last_dts int64
		last_pts int64
	}
	tsRef := make(map[int]*timestampRef)

	// Loop through packets
	for {
		// Read frame
		if err = inputFormatContext.ReadFrame(pkt); err != nil {
			if errors.Is(err, avgo.ErrEof) {
				err = inputFormatContext.SeekFile(
					-1, math.MinInt64, inputFormatContext.StartTime(),
					inputFormatContext.StartTime(), avgo.NewSeekFlags())
				// for i, _ := range outputStreams {
				// }
				continue
				if err != nil {
					log.Fatal(fmt.Errorf("main: seeking frame failed: %w", err))
					break
				} else {
					pktTs, ok := tsRef[pkt.StreamIndex()]
					if ok {
						pktTs.last_dts = pktTs.dts
						pktTs.last_pts = pktTs.pts
					}
					// time.Sleep(1 * time.Second)
					// times++
					continue
				}
			}
			log.Fatal(fmt.Errorf("main: reading frame failed: %w", err))
			break
		}
		// if pkt.Pts() < 0 || pkt.Dts() < 0 {
		// 	continue
		// }
		// if pkt.StreamIndex() == 1 {
		// 	fmt.Print(pkt.Dts(), " ", pkt.Pts(), " ", pkt.Size(), "||")
		// }

		// Get input stream
		inputStream, ok := inputStreams[pkt.StreamIndex()]
		if !ok {
			pkt.Unref()
			continue
		}

		// Get output stream
		outputStream, ok := outputStreams[pkt.StreamIndex()]
		if !ok {
			pkt.Unref()
			continue
		}

		// Update packet

		var dtsOffset int64
		var ptsOffset int64
		pktTs, ok := tsRef[pkt.StreamIndex()]
		if ok {
			dtsOffset = pktTs.last_dts
			ptsOffset = pktTs.last_pts
		} else {
			pktTs = &timestampRef{}
		}

		pkt.SetDts(pkt.Dts() + dtsOffset)
		pkt.SetPts(pkt.Pts() + ptsOffset)
		pktTs.dts = pkt.Dts()
		pktTs.pts = pkt.Pts()
		pktTs.last_dts = pkt.Dts()
		pktTs.last_pts = pkt.Pts()

		tsRef[pkt.StreamIndex()] = pktTs

		pkt.SetStreamIndex(outputStream.Index())
		pkt.RescaleTs(inputStream.TimeBase(), outputStream.TimeBase())
		pkt.SetPos(-1)
		// pkt.SetPts(-1)
		// pkt.SetDts(-1)

		// Write frame
		if err = outputFormatContext.WriteInterleavedFrame(pkt); err != nil {
			log.Fatal(fmt.Errorf("main: writing interleaved frame failed: %w", err))
		}
		// css += pkt.Size()
		// if css > sps {
		// 	fmt.Print("reached ", css, " bytes\n")
		// 	// time.Sleep(5 * time.Millisecond)
		// 	css = 0
		// }
	}

	// Write trailer
	if err = outputFormatContext.WriteTrailer(); err != nil {
		log.Fatal(fmt.Errorf("main: writing trailer failed: %w", err))
	}

	// Success
	log.Println("success")
}

func durationMax(tmp_dur, dur int64, tmp_tb, tb avgo.Rational) int64 {
	tb.Den()
	if tmp_dur == 0 {
		return dur
	}
	if dur == 0 {
		return tmp_dur
	}
	if tmp_dur > dur {
		return tmp_dur
	}
	return dur
}
