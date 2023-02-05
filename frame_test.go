package avgo_test

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func videoInputLastVideoFrame() (f *avgo.Frame, err error) {
	if global.frame != nil {
		return global.frame, nil
	}

	var fc *avgo.FormatContext
	if fc, err = videoInputFormatContext(); err != nil {
		err = fmt.Errorf("avgo_test: getting input format context failed: %w", err)
		return
	}

	var cc *avgo.CodecContext
	var cs *avgo.Stream
	for _, s := range fc.Streams() {
		if s.CodecParameters().MediaType() != avgo.MediaTypeVideo {
			continue
		}

		cs = s

		c := avgo.FindDecoder(s.CodecParameters().CodecID())
		if c == nil {
			err = errors.New("avgo_test: no codec")
			return
		}

		cc = avgo.AllocCodecContext(c)
		if cc == nil {
			err = errors.New("avgo_test: no codec context")
			return
		}
		global.closer.Add(cc.Free)

		if err = cs.CodecParameters().ToCodecContext(cc); err != nil {
			err = fmt.Errorf("avgo_test: updating codec context failed: %w", err)
			return
		}

		if err = cc.Open(c, nil); err != nil {
			err = fmt.Errorf("avgo_test: opening codec context failed: %w", err)
			return
		}
		break
	}

	if cs == nil {
		err = errors.New("avgo_test: no valid video stream")
		return
	}

	var pkt1 *avgo.Packet
	if pkt1, err = videoInputFirstPacket(); err != nil {
		err = fmt.Errorf("avgo_test: getting input first packet failed: %w", err)
		return
	}

	pkt2 := avgo.AllocPacket()
	global.closer.Add(pkt2.Free)

	f = avgo.AllocFrame()
	global.closer.Add(f.Free)

	lastFrame := avgo.AllocFrame()
	global.closer.Add(lastFrame.Free)

	pkts := []*avgo.Packet{pkt1}
	for {
		if err = fc.ReadFrame(pkt2); err != nil {
			if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
				if err = f.Ref(lastFrame); err != nil {
					err = fmt.Errorf("avgo_test: refing frame failed: %w", err)
					return
				}
				err = nil
				break
			}
			err = fmt.Errorf("avgo_test: reading frame failed: %w", err)
			return
		}

		pkts = append(pkts, pkt2)

		for _, pkt := range pkts {
			if pkt.StreamIndex() != cs.Index() {
				continue
			}

			if err = cc.SendPacket(pkt); err != nil {
				err = fmt.Errorf("avgo_test: sending packet failed: %w", err)
				return
			}

			for {
				if err = cc.ReceiveFrame(f); err != nil {
					if errors.Is(err, avgo.ErrEof) || errors.Is(err, avgo.ErrEagain) {
						err = nil
						break
					}
					err = fmt.Errorf("avgo_test: receiving frame failed: %w", err)
					return
				}

				if err = lastFrame.Ref(f); err != nil {
					err = fmt.Errorf("avgo_test: refing frame failed: %w", err)
					return
				}
			}
		}

		pkts = []*avgo.Packet{}
	}
	return
}

func TestFrame(t *testing.T) {
	f1, err := videoInputLastVideoFrame()
	require.NoError(t, err)
	b, err := ioutil.ReadFile("testdata/frame")
	require.NoError(t, err)
	require.Equal(t, string(b), fmt.Sprintf("%+v", f1.Data()))
	require.Equal(t, [8]int{384, 192, 192, 0, 0, 0, 0, 0}, f1.Linesize())
	require.Equal(t, int64(60928), f1.PktDts())
	require.Equal(t, int64(60928), f1.PktPts())

	f2 := avgo.AllocFrame()
	require.NotNil(t, f2)
	defer f2.Free()
	f2.SetChannelLayout(avgo.ChannelLayout21)
	f2.SetHeight(2)
	f2.SetKeyFrame(true)
	f2.SetNbSamples(4)
	f2.SetPictureType(avgo.PictureTypeB)
	f2.SetPixelFormat(avgo.PixelFormat0Bgr)
	require.Equal(t, avgo.PixelFormat0Bgr, f2.PixelFormat()) // Need to test it right away as sample format actually updates the same field
	f2.SetPts(7)
	f2.SetSampleFormat(avgo.SampleFormatDbl)
	require.Equal(t, avgo.SampleFormatDbl, f2.SampleFormat())
	f2.SetSampleRate(9)
	f2.SetWidth(10)
	require.Equal(t, avgo.ChannelLayout21, f2.ChannelLayout())
	require.Equal(t, 2, f2.Height())
	require.True(t, f2.KeyFrame())
	require.Equal(t, 4, f2.NbSamples())
	require.Equal(t, avgo.PictureTypeB, f2.PictureType())
	require.Equal(t, int64(7), f2.Pts())
	require.Equal(t, 9, f2.SampleRate())
	require.Equal(t, 10, f2.Width())

	f3 := f1.Clone()
	require.NotNil(t, f3)
	defer f3.Free()
	require.Equal(t, 180, f3.Height())

	err = f2.AllocBuffer(0)
	require.NoError(t, err)
	err = f3.Ref(f2)
	require.NoError(t, err)
	require.Equal(t, 2, f3.Height())

	f3.MoveRef(f1)
	require.Equal(t, 180, f3.Height())
	require.Equal(t, 0, f1.Height())

	f3.Unref()
	require.Equal(t, 0, f3.Height())

	f4 := avgo.AllocFrame()
	require.NotNil(t, f4)
	defer f4.Free()
	f4.SetNbSamples(960)
	f4.SetChannelLayout(avgo.ChannelLayoutStereo)
	f4.SetSampleFormat(avgo.SampleFormatS16)
	f4.SetSampleRate(48000)
	err = f4.AllocBuffer(0)
	require.NoError(t, err)
	err = f4.AllocSamples(0)
	require.NoError(t, err)

	f5 := avgo.AllocFrame()
	require.NotNil(t, f5)
	defer f5.Free()
	sd := f5.NewSideData(avgo.FrameSideDataTypeAudioServiceType, 4)
	require.NotNil(t, sd)
	sd.SetData([]byte{1, 2, 3})
	sd = f5.SideData(avgo.FrameSideDataTypeAudioServiceType)
	require.NotNil(t, sd)
	require.Equal(t, avgo.FrameSideDataTypeAudioServiceType, sd.Type())
	require.True(t, bytes.HasPrefix(sd.Data(), []byte{1, 2, 3}))
	require.Len(t, sd.Data(), 4)
	sd.SetData([]byte{1, 2, 3, 4, 5})
	sd = f5.SideData(avgo.FrameSideDataTypeAudioServiceType)
	require.NotNil(t, sd)
	require.Equal(t, []byte{1, 2, 3, 4}, sd.Data())

	f6 := avgo.AllocFrame()
	require.NotNil(t, f6)
	defer f6.Free()
	f6.SetHeight(2)
	f6.SetPixelFormat(avgo.PixelFormatYuv420P)
	f6.SetWidth(4)
	require.NoError(t, f6.AllocBuffer(0))
	require.NoError(t, f6.AllocImage(0))
}
