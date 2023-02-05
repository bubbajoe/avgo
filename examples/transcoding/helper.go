package main

// import (
// 	"errors"
// 	"fmt"
// 	"strconv"

// 	"github.com/bubbajoe/avgo"
// )

// func encodeWriteFrame(f *avgo.Frame, s *stream) (err error) {
// 	// Unref packet
// 	s.encPkt.Unref()

// 	// Send frame
// 	if err = s.encCodecContext.SendFrame(f); err != nil {
// 		err = fmt.Errorf("main: sending frame failed: %w", err)
// 		return
// 	}

// 	// Loop
// 	for {
// 		// Receive packet
// 		if err = s.encCodecContext.ReceivePacket(s.encPkt); err != nil {
// 			if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
// 				err = nil
// 				break
// 			}
// 			err = fmt.Errorf("main: receiving packet failed: %w", err)
// 			return
// 		}

// 		// Update pkt
// 		s.encPkt.SetStreamIndex(s.outputStream.Index())
// 		s.encPkt.RescaleTs(s.encCodecContext.TimeBase(), s.outputStream.TimeBase())

// 		// Write frame
// 		if err = outputFormatContext.WriteInterleavedFrame(s.encPkt); err != nil {
// 			err = fmt.Errorf("main: writing frame failed: %w", err)
// 			return
// 		}
// 	}
// 	return
// }

// func filterEncodeWriteFrame(f *avgo.Frame, s *stream) (err error) {
// 	// Add frame
// 	if err = s.buffersrcContext.BuffersrcAddFrame(f, avgo.NewBuffersrcFlags(avgo.BuffersrcFlagKeepRef)); err != nil {
// 		err = fmt.Errorf("main: adding frame failed: %w", err)
// 		return
// 	}

// 	// Loop
// 	for {
// 		// Unref frame
// 		s.filterFrame.Unref()

// 		// Get frame
// 		if err = s.buffersinkContext.BuffersinkGetFrame(s.filterFrame, avgo.NewBuffersinkFlags()); err != nil {
// 			if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
// 				err = nil
// 				break
// 			}
// 			err = fmt.Errorf("main: getting frame failed: %w", err)
// 			return
// 		}

// 		// Reset picture type
// 		s.filterFrame.SetPictureType(avgo.PictureTypeNone)

// 		// Encode and write frame
// 		if err = encodeWriteFrame(s.filterFrame, s); err != nil {
// 			err = fmt.Errorf("main: encoding and writing frame failed: %w", err)
// 			return
// 		}
// 	}
// 	return
// }

// func initFilters() (err error) {
// 	// Loop through output streams
// 	for _, s := range streams {
// 		// Alloc graph
// 		if s.filterGraph = avgo.AllocFilterGraph(); s.filterGraph == nil {
// 			err = errors.New("main: graph is nil")
// 			return
// 		}
// 		c.Add(s.filterGraph.Free)

// 		// Alloc outputs
// 		outputs := avgo.AllocFilterInOut()
// 		if outputs == nil {
// 			err = errors.New("main: outputs is nil")
// 			return
// 		}
// 		c.Add(outputs.Free)

// 		// Alloc inputs
// 		inputs := avgo.AllocFilterInOut()
// 		if inputs == nil {
// 			err = errors.New("main: inputs is nil")
// 			return
// 		}
// 		c.Add(inputs.Free)

// 		// Switch on media type
// 		var args avgo.FilterArgs
// 		var buffersrc, buffersink *avgo.Filter
// 		var content string
// 		switch s.decCodecContext.MediaType() {
// 		case avgo.MediaTypeAudio:
// 			args = avgo.FilterArgs{
// 				"channel_layout": s.decCodecContext.ChannelLayout().String(),
// 				"sample_fmt":     s.decCodecContext.SampleFormat().Name(),
// 				"sample_rate":    strconv.Itoa(s.decCodecContext.SampleRate()),
// 				"time_base":      s.decCodecContext.TimeBase().String(),
// 			}
// 			buffersrc = avgo.FindFilterByName("abuffer")
// 			buffersink = avgo.FindFilterByName("abuffersink")
// 			content = fmt.Sprintf("aformat=sample_fmts=%s:channel_layouts=%s", s.encCodecContext.SampleFormat().Name(), s.encCodecContext.ChannelLayout().String())
// 		default:
// 			args = avgo.FilterArgs{
// 				"pix_fmt":      strconv.Itoa(int(s.decCodecContext.PixelFormat())),
// 				"pixel_aspect": s.decCodecContext.SampleAspectRatio().String(),
// 				"time_base":    s.decCodecContext.TimeBase().String(),
// 				"video_size":   strconv.Itoa(s.decCodecContext.Width()) + "x" + strconv.Itoa(s.decCodecContext.Height()),
// 			}
// 			buffersrc = avgo.FindFilterByName("buffer")
// 			buffersink = avgo.FindFilterByName("buffersink")
// 			content = fmt.Sprintf("format=pix_fmts=%s", s.encCodecContext.PixelFormat().Name())
// 		}

// 		// Check filters
// 		if buffersrc == nil {
// 			err = errors.New("main: buffersrc is nil")
// 			return
// 		}
// 		if buffersink == nil {
// 			err = errors.New("main: buffersink is nil")
// 			return
// 		}

// 		// Create filter contexts
// 		if s.buffersrcContext, err = s.filterGraph.NewFilterContext(buffersrc, "in", args); err != nil {
// 			err = fmt.Errorf("main: creating buffersrc context failed: %w", err)
// 			return
// 		}
// 		if s.buffersinkContext, err = s.filterGraph.NewFilterContext(buffersink, "in", nil); err != nil {
// 			err = fmt.Errorf("main: creating buffersink context failed: %w", err)
// 			return
// 		}

// 		// Update outputs
// 		outputs.SetName("in")
// 		outputs.SetFilterContext(s.buffersrcContext)
// 		outputs.SetPadIdx(0)
// 		outputs.SetNext(nil)

// 		// Update inputs
// 		inputs.SetName("out")
// 		inputs.SetFilterContext(s.buffersinkContext)
// 		inputs.SetPadIdx(0)
// 		inputs.SetNext(nil)

// 		// Parse
// 		if err = s.filterGraph.Parse(content, inputs, outputs); err != nil {
// 			err = fmt.Errorf("main: parsing filter failed: %w", err)
// 			return
// 		}

// 		// Configure
// 		if err = s.filterGraph.Configure(); err != nil {
// 			err = fmt.Errorf("main: configuring filter failed: %w", err)
// 			return
// 		}

// 		// Alloc frame
// 		s.filterFrame = avgo.AllocFrame()
// 		c.Add(s.filterFrame.Free)

// 		// Alloc packet
// 		s.encPkt = avgo.AllocPacket()
// 		c.Add(s.encPkt.Free)
// 	}
// 	return
// }
