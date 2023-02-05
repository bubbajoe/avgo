package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bubbajoe/avgo"
)

var (
	input  = flag.String("i", "", "the input path")
	output = flag.String("o", "", "the output path")
)

var (
	c                   = avgo.NewCloser()
	inputFormatContext  *avgo.FormatContext
	outputFormatContext *avgo.FormatContext
	streams             = make(map[int]*stream) // Indexed by input stream index
)

type stream struct {
	buffersinkContext *avgo.FilterContext
	buffersrcContext  *avgo.FilterContext
	drawtextContext   *avgo.FilterContext
	decCodec          *avgo.Codec
	decCodecContext   *avgo.CodecContext
	decFrame          *avgo.Frame
	encCodec          *avgo.Codec
	encCodecContext   *avgo.CodecContext
	encPkt            *avgo.Packet
	filterFrame       *avgo.Frame
	filterGraph       *avgo.FilterGraph
	inputStream       *avgo.Stream
	outputStream      *avgo.Stream
	startTime         int64
}

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

	defer c.Close()

	// Open input file
	if err := openInputFile(); err != nil {
		log.Fatal(fmt.Errorf("main: opening input file failed: %w", err))
	}

	// Open output file
	if err := openOutputFile(); err != nil {
		log.Fatal(fmt.Errorf("main: opening output file failed: %w", err))
	}

	// Init filters
	if err := initFilters(); err != nil {
		log.Fatal(fmt.Errorf("main: initializing filters failed: %w", err))
	}

	// Alloc packet
	pkt := avgo.AllocPacket()
	c.Add(pkt.Free)
	// Loop through packets
	for {
		// Read frame
		if err := inputFormatContext.ReadFrame(pkt); err != nil {
			if errors.Is(err, avgo.ErrEof) {
				break
			}
			log.Fatal(fmt.Errorf("main: reading frame failed: %w", err))
		}

		// Get stream
		s, ok := streams[pkt.StreamIndex()]
		if !ok {
			continue
		}

		// Update packet
		pkt.RescaleTs(s.inputStream.TimeBase(), s.decCodecContext.TimeBase())

		// if pkt.Dts() > 0 {
		// 	videoTimestamp := int64(float64(pkt.Dts()) * s.decCodecContext.TimeBase().ToDouble() * 1000)
		// 	sendTime := s.startTime + videoTimestamp
		// 	for time.Now().UnixMilli() < sendTime {
		// 		videoTimestamp += 1
		// 	}
		// 	videoTimestamp -= 1
		// }

		// Send packet
		if err := s.decCodecContext.SendPacket(pkt); err != nil {
			log.Fatal(fmt.Errorf("main: sending packet failed: %w", err))
		}

		// Loop
		for {
			// Receive frame
			if err := s.decCodecContext.ReceiveFrame(s.decFrame); err != nil {
				if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
					break
				}
				log.Fatal(fmt.Errorf("main: receiving frame failed: %w", err))
			}

			// Filter, encode and write frame
			if err := filterEncodeWriteFrame(s.decFrame, s); err != nil {
				log.Fatal(fmt.Errorf("main: filtering, encoding and writing frame failed: %w", err))
			}
		}
	}

	// Loop through streams
	for _, s := range streams {
		// Flush filter
		if err := filterEncodeWriteFrame(nil, s); err != nil {
			log.Fatal(fmt.Errorf("main: filtering, encoding and writing frame failed: %w", err))
		}

		// Flush encoder
		if err := encodeWriteFrame(nil, s); err != nil {
			log.Fatal(fmt.Errorf("main: encoding and writing frame failed: %w", err))
		}
	}

	// Write trailer
	if err := outputFormatContext.WriteTrailer(); err != nil {
		log.Fatal(fmt.Errorf("main: writing trailer failed: %w", err))
	}

	// Success
	log.Println("success")
}

func openInputFile() (err error) {
	// Alloc input format context
	if inputFormatContext = avgo.AllocFormatContext(); inputFormatContext == nil {
		err = errors.New("main: input format context is nil")
		return
	}
	c.Add(inputFormatContext.Free)
	{
		dict := avgo.NewDictionary()
		defer dict.Free()
		dict.Set("safe", "0", avgo.DictionaryFlags(avgo.DictionaryFlagAppend))
		// dict.Set("re", "1", avgo.DictionaryFlags(avgo.DictionaryFlagAppend))
		// dict.Set("stream_loop", "-1", avgo.DictionaryFlags(avgo.DictionaryFlagAppend))
		// Open input
		if err = inputFormatContext.OpenInput(*input, nil, dict); err != nil {
			err = fmt.Errorf("main: opening input failed: %w", err)
			return
		}
	}
	c.Add(inputFormatContext.CloseInput)

	// Find stream info
	if err = inputFormatContext.FindStreamInfo(nil); err != nil {
		err = fmt.Errorf("main: finding stream info failed: %w", err)
		return
	}

	// Loop through streams
	for _, is := range inputFormatContext.Streams() {
		// Only process audio or video
		if is.CodecParameters().MediaType() != avgo.MediaTypeAudio &&
			is.CodecParameters().MediaType() != avgo.MediaTypeVideo {
			continue
		}

		// Create stream
		s := &stream{inputStream: is}

		s.startTime = time.Now().UnixMilli()

		// Find decoder
		if s.decCodec = avgo.FindDecoder(is.CodecParameters().CodecID()); s.decCodec == nil {
			err = errors.New("main: codec is nil")
			return
		}

		// Alloc codec context
		if s.decCodecContext = avgo.AllocCodecContext(s.decCodec); s.decCodecContext == nil {
			err = errors.New("main: codec context is nil")
			return
		}
		c.Add(s.decCodecContext.Free)

		// Update codec context
		if err = is.CodecParameters().ToCodecContext(s.decCodecContext); err != nil {
			err = fmt.Errorf("main: updating codec context failed: %w", err)
			return
		}

		// Set framerate
		if is.CodecParameters().MediaType() == avgo.MediaTypeVideo {
			s.decCodecContext.SetFramerate(inputFormatContext.GuessFrameRate(is, nil))
		}

		// Open codec context
		if err = s.decCodecContext.Open(s.decCodec, nil); err != nil {
			err = fmt.Errorf("main: opening codec context failed: %w", err)
			return
		}

		// Alloc frame
		s.decFrame = avgo.AllocFrame()
		c.Add(s.decFrame.Free)

		// Store stream
		streams[is.Index()] = s
	}
	return
}

func openOutputFile() (err error) {
	format := ""
	// Alloc output format context
	if strings.HasPrefix(*output, "rtmp://") {
		format = "flv"
	}
	if outputFormatContext, err = avgo.AllocOutputFormatContext(nil, format, *output); err != nil {
		err = fmt.Errorf("main: allocating output format context failed: %w", err)
		return
	} else if outputFormatContext == nil {
		err = errors.New("main: output format context is nil")
		return
	}
	c.Add(outputFormatContext.Free)

	// Loop through streams
	for _, is := range inputFormatContext.Streams() {
		// Get stream
		s, ok := streams[is.Index()]
		if !ok {
			continue
		}

		// Create output stream
		if s.outputStream = outputFormatContext.NewStream(nil); s.outputStream == nil {
			err = errors.New("main: output stream is nil")
			return
		}

		// Get codec id
		codecID := avgo.CodecIDH264
		if s.decCodecContext.MediaType() == avgo.MediaTypeAudio {
			codecID = avgo.CodecIDAac
		}

		// Find encoder
		if s.encCodec = avgo.FindEncoder(codecID); s.encCodec == nil {
			err = errors.New("main: codec is nil")
			return
		}

		// Alloc codec context
		if s.encCodecContext = avgo.AllocCodecContext(s.encCodec); s.encCodecContext == nil {
			err = errors.New("main: codec context is nil")
			return
		}
		c.Add(s.encCodecContext.Free)

		// Update codec context
		if s.decCodecContext.MediaType() == avgo.MediaTypeAudio {
			if v := s.encCodec.ChannelLayouts(); len(v) > 0 {
				s.encCodecContext.SetChannelLayout(v[0])
			} else {
				s.encCodecContext.SetChannelLayout(s.decCodecContext.ChannelLayout())
			}
			s.encCodecContext.SetChannels(s.decCodecContext.Channels())
			s.encCodecContext.SetSampleRate(s.decCodecContext.SampleRate())
			if v := s.encCodec.SampleFormats(); len(v) > 0 {
				s.encCodecContext.SetSampleFormat(v[0])
			} else {
				s.encCodecContext.SetSampleFormat(s.decCodecContext.SampleFormat())
			}
			s.encCodecContext.SetTimeBase(s.decCodecContext.TimeBase())
			// } else if s.decCodecContext.MediaType() == avgo.MediaTypeVideo {
		} else {
			s.encCodecContext.SetHeight(s.decCodecContext.Height())
			if v := s.encCodec.PixelFormats(); len(v) > 0 {
				s.encCodecContext.SetPixelFormat(v[0])
			} else {
				s.encCodecContext.SetPixelFormat(s.decCodecContext.PixelFormat())
			}
			s.encCodecContext.SetSampleAspectRatio(s.decCodecContext.SampleAspectRatio())
			s.encCodecContext.SetTimeBase(s.decCodecContext.TimeBase())
			s.encCodecContext.SetWidth(s.decCodecContext.Width())
		}

		// Update flags
		if s.decCodecContext.Flags().Has(avgo.CodecContextFlagGlobalHeader) {
			s.encCodecContext.SetFlags(s.encCodecContext.Flags().Add(avgo.CodecContextFlagGlobalHeader))
		}

		// Open codec context
		if err = s.encCodecContext.Open(s.encCodec, nil); err != nil {
			err = fmt.Errorf("main: opening codec context failed: %w", err)
			return
		}

		// Update codec parameters
		if err = s.outputStream.CodecParameters().FromCodecContext(s.encCodecContext); err != nil {
			err = fmt.Errorf("main: updating codec parameters failed: %w", err)
			return
		}

		// Update stream
		s.outputStream.SetTimeBase(s.encCodecContext.TimeBase())
	}

	// If this is a file, we need to use an io context
	if !outputFormatContext.OutputFormat().Flags().Has(avgo.IOFormatFlagNofile) {
		// Create io context
		ioContext := avgo.NewIOContext()
		{
			dict := avgo.NewDictionary()
			defer dict.Free()
			// Open io context
			if err = ioContext.OpenWith(*output, avgo.NewIOContextFlags(avgo.IOContextFlagWrite), dict); err != nil {
				err = fmt.Errorf("main: opening io context failed: %w", err)
				return
			}
		}
		c.AddWithError(ioContext.Closep)

		// Update output format context
		outputFormatContext.SetPb(ioContext)
	}

	// Write header
	if err = outputFormatContext.WriteHeader(nil); err != nil {
		err = fmt.Errorf("main: writing header failed: %w", err)
		return
	}
	return
}

func initFilters() (err error) {
	// Loop through output streams
	for _, s := range streams {
		// Alloc graph
		if s.filterGraph = avgo.AllocFilterGraph(); s.filterGraph == nil {
			err = errors.New("main: graph is nil")
			return
		}
		c.Add(s.filterGraph.Free)

		// Alloc outputs
		outputs := avgo.AllocFilterInOut()
		if outputs == nil {
			err = errors.New("main: outputs is nil")
			return
		}
		c.Add(outputs.Free)

		// Alloc inputs
		inputs := avgo.AllocFilterInOut()
		if inputs == nil {
			err = errors.New("main: inputs is nil")
			return
		}
		c.Add(inputs.Free)

		// Alloc inputs
		puts := avgo.AllocFilterInOut()
		if puts == nil {
			err = errors.New("main: puts is nil")
			return
		}
		c.Add(puts.Free)

		// Switch on media type
		var args avgo.FilterArgs
		var buffersrc, buffersink, drawfilter *avgo.Filter
		var content string
		switch s.decCodecContext.MediaType() {
		case avgo.MediaTypeAudio:
			args = avgo.FilterArgs{
				"channel_layout": s.decCodecContext.ChannelLayout().String(),
				"sample_fmt":     s.decCodecContext.SampleFormat().Name(),
				"sample_rate":    strconv.Itoa(s.decCodecContext.SampleRate()),
				"time_base":      s.decCodecContext.TimeBase().String(),
			}
			buffersrc = avgo.FindFilterByName("abuffer")
			buffersink = avgo.FindFilterByName("abuffersink")
			content = fmt.Sprintf("aformat=sample_fmts=%s:channel_layouts=%s", s.encCodecContext.SampleFormat().Name(), s.encCodecContext.ChannelLayout().String())
		default:
			args = avgo.FilterArgs{
				"pix_fmt":      strconv.Itoa(int(s.decCodecContext.PixelFormat())),
				"pixel_aspect": s.decCodecContext.SampleAspectRatio().String(),
				"time_base":    s.decCodecContext.TimeBase().String(),
				"video_size":   strconv.Itoa(s.decCodecContext.Width()) + "x" + strconv.Itoa(s.decCodecContext.Height()),
			}
			buffersrc = avgo.FindFilterByName("buffer")
			buffersink = avgo.FindFilterByName("buffersink")
			drawfilter = avgo.FindFilterByName("drawtext")
			content = fmt.Sprintf("format=pix_fmts=%s", s.encCodecContext.PixelFormat().Name())
		}

		// Check filters
		if buffersrc == nil {
			err = errors.New("main: buffersrc is nil")
			return
		}
		if buffersink == nil {
			err = errors.New("main: buffersink is nil")
			return
		}
		if drawfilter == nil {
			err = errors.New("main: drawfilter is nil")
			return
		}

		// Create filter contexts
		if s.buffersrcContext, err = s.filterGraph.NewFilterContext(buffersrc, "1", args); err != nil {
			err = fmt.Errorf("main: creating buffersrc context failed: %w", err)
			return
		}
		if s.buffersinkContext, err = s.filterGraph.NewFilterContext(buffersink, "1", nil); err != nil {
			err = fmt.Errorf("main: creating buffersink context failed: %w", err)
			return
		}
		if s.drawtextContext, err = s.filterGraph.NewFilterContext(drawfilter, "1", avgo.FilterArgs{
			// "fontfile": "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf",
			"text":      "Hello World",
			"fontsize":  "30",
			"x":         "10",
			"y":         "10",
			"fontcolor": "white",
			"box":       "1",
		}); err != nil {
			err = fmt.Errorf("main: creating drawtext context failed: %w", err)
			return
		}

		// Update outputs
		outputs.SetName("1")
		outputs.SetFilterContext(s.buffersrcContext)
		outputs.SetPadIdx(0)
		outputs.SetNext(nil)

		// Update outputs
		puts.SetName("1")
		puts.SetFilterContext(s.drawtextContext)
		puts.SetPadIdx(0)
		puts.SetNext(nil)

		// Update inputs
		inputs.SetName("1")
		inputs.SetFilterContext(s.buffersinkContext)
		inputs.SetPadIdx(0)
		inputs.SetNext(nil)

		// Parse
		if err = s.filterGraph.Parse(content, inputs, outputs); err != nil {
			err = fmt.Errorf("main: parsing filter failed: %w", err)
			return
		}
		if err = s.filterGraph.Parse(content, puts, outputs); err != nil {
			err = fmt.Errorf("main: parsing filter failed: %w", err)
			return
		}

		// Configure
		if err = s.filterGraph.Configure(); err != nil {
			err = fmt.Errorf("main: configuring filter failed: %w", err)
			return
		}

		// Alloc frame
		s.filterFrame = avgo.AllocFrame()
		c.Add(s.filterFrame.Free)

		// Alloc packet
		s.encPkt = avgo.AllocPacket()
		c.Add(s.encPkt.Free)
	}
	return
}

func filterEncodeWriteFrame(f *avgo.Frame, s *stream) (err error) {
	// Add frame
	if err = s.buffersrcContext.BuffersrcAddFrame(f, avgo.NewBuffersrcFlags(avgo.BuffersrcFlagKeepRef)); err != nil {
		err = fmt.Errorf("main: adding frame failed: %w", err)
		return
	}

	// Loop
	for {
		// Unref frame
		s.filterFrame.Unref()

		// Get frame
		if err = s.buffersinkContext.BuffersinkGetFrame(s.filterFrame, avgo.NewBuffersinkFlags()); err != nil {
			if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
				err = nil
				break
			}
			err = fmt.Errorf("main: getting frame failed: %w", err)
			return
		}

		// Reset picture type
		s.filterFrame.SetPictureType(avgo.PictureTypeNone)

		// Encode and write frame
		if err = encodeWriteFrame(s.filterFrame, s); err != nil {
			err = fmt.Errorf("main: encoding and writing frame failed: %w", err)
			return
		}
	}
	return
}

func encodeWriteFrame(f *avgo.Frame, s *stream) (err error) {
	// Unref packet
	s.encPkt.Unref()

	// Send frame
	if err = s.encCodecContext.SendFrame(f); err != nil {
		err = fmt.Errorf("main: sending frame failed: %w", err)
		return
	}

	// Loop
	for {
		// Receive packet
		if err = s.encCodecContext.ReceivePacket(s.encPkt); err != nil {
			if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
				err = nil
				break
			}
			err = fmt.Errorf("main: receiving packet failed: %w", err)
			return
		}

		// Update pkt
		s.encPkt.SetStreamIndex(s.outputStream.Index())
		s.encPkt.RescaleTs(s.encCodecContext.TimeBase(), s.outputStream.TimeBase())

		tb := s.outputStream.TimeBase()
		if s.encPkt.Pts() > 0 {
			videoTimestamp := int64(float64(s.encPkt.Pts()) * tb.ToDouble() * 1000)
			sendTime := s.startTime + videoTimestamp
			for time.Now().UnixMilli() < sendTime {
				videoTimestamp += 1
			}
			videoTimestamp -= 1
		}

		// Write frame
		if err = outputFormatContext.WriteInterleavedFrame(s.encPkt); err != nil {
			err = fmt.Errorf("main: writing frame failed: %w", err)
			return
		}
	}
	return
}
