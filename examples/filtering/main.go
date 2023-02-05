package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

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
	s                   *stream
	streams             = make(map[int]*stream)
)

type stream struct {
	buffersinkContext *avgo.FilterContext
	buffersrcContext  *avgo.FilterContext
	buffersrc2Context *avgo.FilterContext
	overlayContext    *avgo.FilterContext
	drawtextContext   *avgo.FilterContext
	decCodec          *avgo.Codec
	decCodecContext   *avgo.CodecContext
	decFrame          *avgo.Frame
	encCodec          *avgo.Codec
	encCodecContext   *avgo.CodecContext
	encFrame          *avgo.Frame
	filterFrame       *avgo.Frame
	filterGraph       *avgo.FilterGraph
	inputStream       *avgo.Stream
	outputStream      *avgo.Stream
	lastPts           int64
}

func main() {
	// Handle ffmpeg logs
	avgo.SetLogLevel(avgo.LogLevelDebug)
	avgo.SetLogCallback(func(l avgo.LogLevel, fmt, msg, parent string) {
		log.Printf("ffmpeg log: %s (level: %s/%d)\n", strings.TrimSpace(msg), l, l)
	})

	// Parse flags
	flag.Parse()

	// Usage
	if *input == "" {
		log.Println("Usage: <binary path> -i <input path>")
		return
	}

	defer c.Close()

	streams = make(map[int]*stream)

	// Open input file
	if err := openInputFile(); err != nil {
		log.Fatal(fmt.Errorf("main: opening input file failed: %w", err))
	}

	// Open output file
	if err := openOutputFile(); err != nil {
		log.Fatal(fmt.Errorf("main: opening output file failed: %w", err))
	}

	// Init filter
	if err := initFilter(); err != nil {
		log.Fatal(fmt.Errorf("main: initializing filter failed: %w", err))
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

		// Invalid stream
		if pkt.StreamIndex() != s.inputStream.Index() {
			continue
		}

		// Send packet
		if err := s.decCodecContext.SendPacket(pkt); err != nil {
			log.Fatal(fmt.Errorf("main: sending packet failed: %w", err))
		}

		// Loop
		for {
			// Receive decoded frame
			if err := s.decCodecContext.ReceiveFrame(s.decFrame); err != nil {
				if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
					break
				}
				log.Fatal(fmt.Errorf("main: receiving frame failed: %w", err))
			}

			// Filter decoded frame
			if err := filterFrame(s.decFrame, s); err != nil {
				log.Fatal(fmt.Errorf("main: filtering frame failed: %w", err))
			}
		}
	}

	// Flush filter
	if err := filterFrame(nil, s); err != nil {
		log.Fatal(fmt.Errorf("main: flushing filtering frame failed: %w", err))
	}

	// Flush encoder
	if err := encodeFrame(nil, s); err != nil {
		log.Fatal(fmt.Errorf("main: flushing encoder frame failed: %w", err))
	}

	err := outputFormatContext.WriteTrailer()
	if err != nil {
		log.Fatal(fmt.Errorf("main: writing trailer failed: %w", err))
	}

	// Success
	log.Println("success")
}

func openOutputFile() (err error) {
	// Alloc output format context
	if outputFormatContext, err = avgo.AllocOutputFormatContext(nil, "", *output); err != nil {
		err = fmt.Errorf("main: allocating output format context failed: %w", err)
		return
	} else if outputFormatContext == nil {
		err = errors.New("main: output format context is nil")
		return
	}
	c.Add(outputFormatContext.Free)

	//Loop through streams
	for _, is := range inputFormatContext.Streams() {
		// Get input stream
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
		} else if s.decCodecContext.MediaType() == avgo.MediaTypeVideo {
			s.encCodecContext.SetHeight(s.decCodecContext.Height())
			if v := s.encCodec.PixelFormats(); len(v) > 0 {
				s.encCodecContext.SetPixelFormat(v[0])
			} else {
				s.encCodecContext.SetPixelFormat(s.decCodecContext.PixelFormat())
			}
			s.encCodecContext.SetSampleAspectRatio(s.inputStream.SampleAspectRatio())
			s.encCodecContext.SetTimeBase(s.inputStream.TimeBase())
			s.encCodecContext.SetWidth(s.decCodecContext.Width())
			fmt.Println(s.encCodecContext.TimeBase().String())
		}

		// Update flags
		if s.decCodecContext.Flags().Has(avgo.CodecContextFlagGlobalHeader) {
			s.encCodecContext.SetFlags(s.encCodecContext.Flags().Add(avgo.CodecContextFlagGlobalHeader))
		}

		// Open codec context
		if err = s.encCodecContext.Open(s.encCodec, nil); err != nil {
			err = fmt.Errorf("main: opening encode codec context failed: %w", err)
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
	fmt.Println("Using IOContext")

	// If this is a file, we need to use an io context
	if !outputFormatContext.OutputFormat().Flags().Has(avgo.IOFormatFlagNofile) {
	}

	// Write header
	if err = outputFormatContext.WriteHeader(nil); err != nil {
		err = fmt.Errorf("main: writing header failed: %w", err)
		return
	}
	return
}

func openInputFile() (err error) {
	// Alloc input format context
	if inputFormatContext = avgo.AllocFormatContext(); inputFormatContext == nil {
		err = errors.New("main: input format context is nil")
		return
	}
	c.Add(inputFormatContext.Free)

	// Open input
	if err = inputFormatContext.OpenInput(*input, nil, nil); err != nil {
		err = fmt.Errorf("main: opening input failed: %w", err)
		return
	}
	c.Add(inputFormatContext.CloseInput)

	// Find stream info
	if err = inputFormatContext.FindStreamInfo(nil); err != nil {
		err = fmt.Errorf("main: finding stream info failed: %w", err)
		return
	}

	// Loop through streams
	for _, is := range inputFormatContext.Streams() {
		// Only process video
		if is.CodecParameters().MediaType() != avgo.MediaTypeVideo {
			continue
		}

		// Create stream
		s = &stream{
			inputStream: is,
			lastPts:     avgo.NoPtsValue,
		}

		// Add stream
		streams[is.Index()] = s

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

		// Open codec context
		if err = s.decCodecContext.Open(s.decCodec, nil); err != nil {
			err = fmt.Errorf("main: opening decode codec context failed: %w", err)
			return
		}

		// Alloc frame
		s.decFrame = avgo.AllocFrame()
		c.Add(s.decFrame.Free)

		break
	}

	// No video stream
	if s == nil {
		err = errors.New("main: no video stream")
		return
	}
	return
}

func initFilter() (err error) {
	// Alloc graph
	if s.filterGraph = avgo.AllocFilterGraph(); s.filterGraph == nil {
		err = errors.New("main: graph is nil")
		return
	}
	c.Add(s.filterGraph.Free)

	// // Alloc outputs
	// outputs := avgo.AllocFilterInOut()
	// if outputs == nil {
	// 	err = errors.New("main: outputs is nil")
	// 	return
	// }
	// c.Add(outputs.Free)

	// // Alloc inputs
	// inputs := avgo.AllocFilterInOut()
	// if inputs == nil {
	// 	err = errors.New("main: inputs is nil")
	// 	return
	// }
	// c.Add(inputs.Free)

	// Create buffersrc
	buffersrc := avgo.FindFilterByName("buffer")
	if buffersrc == nil {
		err = errors.New("main: buffersrc is nil")
		return
	}

	// Create buffersrc
	overlay := avgo.FindFilterByName("overlay")
	if buffersrc == nil {
		err = errors.New("main: buffersrc is nil")
		return
	}

	// Create buffersrc
	drawtext := avgo.FindFilterByName("drawtext")
	if buffersrc == nil {
		err = errors.New("main: buffersrc is nil")
		return
	}

	// Create buffersink
	buffersink := avgo.FindFilterByName("buffersink")
	if buffersink == nil {
		err = errors.New("main: buffersink is nil")
		return
	}

	// Create filter contexts
	if s.buffersrcContext, err = s.filterGraph.NewFilterContext(buffersrc, "input", avgo.FilterArgs{
		"pix_fmt":      strconv.Itoa(int(s.decCodecContext.PixelFormat())),
		"pixel_aspect": s.decCodecContext.SampleAspectRatio().String(),
		"time_base":    s.inputStream.TimeBase().String(),
		"video_size":   strconv.Itoa(s.decCodecContext.Width()) + "x" + strconv.Itoa(s.decCodecContext.Height()),
	}); err != nil {
		err = fmt.Errorf("main: creating buffersrc context failed: %w", err)
		return
	}
	if s.buffersrc2Context, err = s.filterGraph.NewFilterContext(buffersrc, "input", avgo.FilterArgs{
		"pix_fmt":      strconv.Itoa(int(s.decCodecContext.PixelFormat())),
		"pixel_aspect": s.decCodecContext.SampleAspectRatio().String(),
		"time_base":    s.inputStream.TimeBase().String(),
		"video_size":   strconv.Itoa(s.decCodecContext.Width()) + "x" + strconv.Itoa(s.decCodecContext.Height()),
	}); err != nil {
		err = fmt.Errorf("main: creating buffersrc context failed: %w", err)
		return
	}
	// x=W-w-%d:y=H-h-%d:repeatlast=1
	if s.overlayContext, err = s.filterGraph.NewFilterContext(overlay, "overlay", avgo.FilterArgs{
		"x":          "W-w+(w-W/2)",
		"y":          "H-h+(h-H/2)",
		"repeatlast": "1",
	}); err != nil {
		err = fmt.Errorf("main: creating overlay context failed: %w", err)
		return
	}
	if s.drawtextContext, err = s.filterGraph.NewFilterContext(drawtext, "text", avgo.FilterArgs{
		"x":        "(w-text_w)/2",
		"y":        "(h-text_h)/2 - 100",
		"text":     "πø s.inputStream.TimeBase().String()",
		"fontsize": "30",
	}); err != nil {
		err = fmt.Errorf("main: creating drawtext context failed: %w", err)
		return
	}
	var scaleCtx *avgo.FilterContext
	if scaleCtx, err = s.filterGraph.NewFilterContextRawArgs(avgo.FindFilterByName("scale"), "scale", "w=iw/3:h=ih/3"); err != nil {
		err = fmt.Errorf("main: creating buffersink context failed: %w", err)
		return
	}
	if s.buffersinkContext, err = s.filterGraph.NewFilterContext(buffersink, "out", nil); err != nil {
		err = fmt.Errorf("main: creating buffersink context failed: %w", err)
		return
	}

	err = s.buffersrcContext.Link(0, s.drawtextContext, 0)
	if err != nil {
		fmt.Println(err)
	}
	err = s.buffersrc2Context.Link(0, scaleCtx, 0)
	if err != nil {
		fmt.Println(err)
	}
	err = s.drawtextContext.Link(0, s.overlayContext, 0)
	if err != nil {
		fmt.Println(err)
	}
	err = scaleCtx.Link(0, s.overlayContext, 1)
	if err != nil {
		fmt.Println(err)
	}
	err = s.overlayContext.Link(0, s.buffersinkContext, 0)
	if err != nil {
		fmt.Println(err)
	}

	// // Update outputs
	// outputs.SetName("in")
	// outputs.SetFilterContext(s.buffersrcContext)
	// outputs.SetPadIdx(0)
	// outputs.SetNext(nil)

	// // Update inputs
	// inputs.SetName("out")
	// inputs.SetFilterContext(s.buffersinkContext)
	// inputs.SetPadIdx(0)
	// inputs.SetNext(nil)

	// // Parse
	// if err = s.filterGraph.Parse("transpose=cclock", inputs, outputs); err != nil {
	// 	err = fmt.Errorf("main: parsing filter failed: %w", err)
	// 	return
	// }

	// Configure
	if err = s.filterGraph.Configure(); err != nil {
		err = fmt.Errorf("main: configuring filter failed: %w", err)
		return
	}

	// Alloc frame
	s.filterFrame = avgo.AllocFrame()
	c.Add(s.filterFrame.Free)
	return
}

func filterFrame(f *avgo.Frame, s *stream) (err error) {
	// Add frame
	if err = s.buffersrcContext.BuffersrcAddFrame(f, avgo.NewBuffersrcFlags(avgo.BuffersrcFlagKeepRef)); err != nil {
		err = fmt.Errorf("main: adding frame failed: %w", err)
		return
	}
	if err = s.buffersrc2Context.BuffersrcAddFrame(f, avgo.NewBuffersrcFlags(avgo.BuffersrcFlagKeepRef)); err != nil {
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

		// Do something with filtered frame
		log.Printf("new filtered frame: %dx%d\n", s.filterFrame.Width(), s.filterFrame.Height())

		// Encode filtered frame
		if err := encodeFrame(s.filterFrame, s); err != nil {
			log.Fatal(fmt.Errorf("main: encoding frame failed: %w", err))
		}
	}
	return
}

func encodeFrame(f *avgo.Frame, s *stream) (err error) {
	// Unref packet
	pkt := avgo.AllocPacket()
	defer pkt.Unref()

	// Send frame
	if err = s.encCodecContext.SendFrame(f); err != nil {
		err = fmt.Errorf("main: sending frame failed: %w", err)
		return
	}

	// Loop
	for {
		// Receive packet
		if err = s.encCodecContext.ReceivePacket(pkt); err != nil {
			if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
				err = nil
				break
			}
			err = fmt.Errorf("main: receiving packet failed: %w", err)
			return
		}

		// Update pkt
		pkt.SetStreamIndex(s.inputStream.Index())
		pkt.RescaleTs(s.encCodecContext.TimeBase(), s.decCodecContext.TimeBase())

		// Write frame
		if err = outputFormatContext.WriteInterleavedFrame(pkt); err != nil {
			err = fmt.Errorf("main: writing frame failed: %w", err)
			return
		}
	}

	return
}
