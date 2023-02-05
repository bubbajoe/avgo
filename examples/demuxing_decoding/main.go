package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/bubbajoe/avgo"
)

var (
	input = flag.String("i", "", "the input path")
)

type stream struct {
	decCodec        *avgo.Codec
	decCodecContext *avgo.CodecContext
	inputStream     *avgo.Stream
}

func main() {
	// Handle ffmpeg logs
	avgo.SetLogLevel(avgo.LogLevelDebug)
	avgo.SetLogCallback(func(l avgo.LogLevel, fmt, msg, parent string) {
		log.Printf("ffmpeg log: %s (level: %d)\n", strings.TrimSpace(msg), l)
	})

	// Parse flags
	flag.Parse()

	// Usage
	if *input == "" {
		log.Println("Usage: <binary path> -i <input path>")
		return
	}

	// Alloc packet
	pkt := avgo.AllocPacket()
	defer pkt.Free()

	// Alloc frame
	f := avgo.AllocFrame()
	defer f.Free()

	// Alloc input format context
	inputFormatContext := avgo.AllocFormatContext()
	if inputFormatContext == nil {
		log.Fatal(errors.New("main: input format context is nil"))
	}
	defer inputFormatContext.Free()

	// Open input
	if err := inputFormatContext.OpenInput(*input, nil, nil); err != nil {
		log.Fatal(fmt.Errorf("main: opening input failed: %w", err))
	}
	defer inputFormatContext.CloseInput()

	// Find stream info
	if err := inputFormatContext.FindStreamInfo(nil); err != nil {
		log.Fatal(fmt.Errorf("main: finding stream info failed: %w", err))
	}

	// Loop through streams
	streams := make(map[int]*stream) // Indexed by input stream index
	for _, is := range inputFormatContext.Streams() {
		// Only process audio or video
		if is.CodecParameters().MediaType() != avgo.MediaTypeAudio &&
			is.CodecParameters().MediaType() != avgo.MediaTypeVideo {
			continue
		}

		// Create stream
		s := &stream{inputStream: is}

		// Find decoder
		if s.decCodec = avgo.FindDecoder(is.CodecParameters().CodecID()); s.decCodec == nil {
			log.Fatal(errors.New("main: codec is nil"))
		}

		// Alloc codec context
		if s.decCodecContext = avgo.AllocCodecContext(s.decCodec); s.decCodecContext == nil {
			log.Fatal(errors.New("main: codec context is nil"))
		}
		defer s.decCodecContext.Free()

		// Update codec context
		if err := is.CodecParameters().ToCodecContext(s.decCodecContext); err != nil {
			log.Fatal(fmt.Errorf("main: updating codec context failed: %w", err))
		}

		// Open codec context
		if err := s.decCodecContext.Open(s.decCodec, nil); err != nil {
			log.Fatal(fmt.Errorf("main: opening codec context failed: %w", err))
		}

		// Add stream
		streams[is.Index()] = s
	}

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

		// Send packet
		if err := s.decCodecContext.SendPacket(pkt); err != nil {
			log.Fatal(fmt.Errorf("main: sending packet failed: %w", err))
		}

		// Loop
		for {
			// Receive frame
			if err := s.decCodecContext.ReceiveFrame(f); err != nil {
				if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
					break
				}
				log.Fatal(fmt.Errorf("main: receiving frame failed: %w", err))
			}

			// Do something with decoded frame
			log.Printf("new frame: stream %d - pts: %d", pkt.StreamIndex(), f.Pts())
		}
	}

	// Success
	log.Println("success")
}
