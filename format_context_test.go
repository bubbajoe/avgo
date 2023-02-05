package avgo_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func videoInputFormatContext() (fc1 *avgo.FormatContext, err error) {
	if global.inputFormatContext != nil {
		return global.inputFormatContext, nil
	}

	if fc1 = avgo.AllocFormatContext(); fc1 == nil {
		err = errors.New("avgo_test: allocated format context is nil")
		return
	}
	global.closer.Add(fc1.Free)

	if err = fc1.OpenInput("testdata/video.mp4", nil, nil); err != nil {
		err = fmt.Errorf("avgo_test: opening input failed: %w", err)
		return
	}
	global.closer.Add(fc1.CloseInput)

	if err = fc1.FindStreamInfo(nil); err != nil {
		err = fmt.Errorf("avgo_test: finding stream info failed: %w", err)
		return
	}

	global.inputFormatContext = fc1
	return
}

func TestFormatContext(t *testing.T) {
	fc1, s1, _, err := videoInputStreams()
	require.NoError(t, err)

	require.Equal(t, int64(607583), fc1.BitRate())
	require.Equal(t, avgo.NewFormatContextCtxFlags(0), fc1.CtxFlags())
	require.Equal(t, int64(5014000), fc1.Duration())
	require.True(t, fc1.EventFlags().Has(avgo.FormatEventFlagMetadataUpdated))
	require.Equal(t, "testdata/video.mp4", fc1.Filename())
	require.True(t, fc1.Flags().Has(avgo.FormatContextFlagAutoBsf))
	require.Equal(t, avgo.NewRational(24, 1), fc1.GuessFrameRate(s1, nil))
	require.Equal(t, avgo.NewRational(1, 1), fc1.GuessSampleAspectRatio(s1, nil))
	require.True(t, fc1.InputFormat().Flags().Has(avgo.IOFormatFlagNoByteSeek))
	require.Equal(t, avgo.IOContextFlags(0), fc1.IOFlags())
	require.Equal(t, int64(0), fc1.MaxAnalyzeDuration())
	require.Equal(t, "isom", fc1.Metadata().Get("major_brand", nil, avgo.NewDictionaryFlags()).Value())
	require.Equal(t, int64(0), fc1.StartTime())
	require.Equal(t, 2, fc1.NbStreams())
	require.Len(t, fc1.Streams(), 2)

	sdp, err := fc1.SDPCreate()
	require.NoError(t, err)
	require.Equal(t, "v=0\r\no=- 0 0 IN IP4 127.0.0.1\r\ns=Big Buck Bunny\r\nt=0 0\r\na=tool:libavformat 58.76.100\r\nm=video 0 RTP/AVP 96\r\nb=AS:441\r\na=rtpmap:96 H264/90000\r\na=fmtp:96 packetization-mode=1; sprop-parameter-sets=Z0LADasgKDPz4CIAAAMAAgAAAwBhHihUkA==,aM48gA==; profile-level-id=42C00D\r\na=control:streamid=0\r\nm=audio 0 RTP/AVP 97\r\nb=AS:161\r\na=rtpmap:97 MPEG4-GENERIC/48000/2\r\na=fmtp:97 profile-level-id=1;mode=AAC-hbr;sizelength=13;indexlength=3;indexdeltalength=3; config=1190\r\na=control:streamid=1\r\n", sdp)

	fc2, err := avgo.AllocOutputFormatContext(nil, "", "/tmp/test.mp4")
	require.NoError(t, err)
	defer fc2.Free()
	require.Equal(t, "/tmp/test.mp4", fc2.Filename())
	require.True(t, fc2.OutputFormat().Flags().Has(avgo.IOFormatFlagGlobalheader))

	fc3 := avgo.AllocFormatContext()
	require.NotNil(t, fc3)
	defer fc3.Free()
	c := avgo.NewIOContext()
	err = c.Open("testdata/video.mp4", avgo.NewIOContextFlags(avgo.IOContextFlagRead))
	require.NoError(t, err)
	defer c.Closep() //nolint:errcheck
	fc3.SetPb(c)
	fc3.SetStrictStdCompliance(avgo.StrictStdComplianceExperimental)
	require.NotNil(t, fc3.Pb())
	require.Equal(t, avgo.StrictStdComplianceExperimental, fc3.StrictStdCompliance())
	s2 := fc3.NewStream(nil)
	require.NotNil(t, s2)
	s3 := fc3.NewStream(nil)
	require.NotNil(t, s3)
	require.Equal(t, 1, s3.Index())

	// TODO Test SetInterruptCallback
	// TODO Test ReadFrame
	// TODO Test SeekFrame
	// TODO Test Flush
	// TODO Test WriteHeader
	// TODO Test WriteFrame
	// TODO Test WriteInterleavedFrame
	// TODO Test WriteTrailer
}
