package avgo_test

import (
	"fmt"
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func videoInputStreams() (fc *avgo.FormatContext, s1, s2 *avgo.Stream, err error) {
	if global.inputFormatContext != nil && global.inputStream1 != nil && global.inputStream2 != nil {
		return global.inputFormatContext, global.inputStream1, global.inputStream2, nil
	}

	if fc, err = videoInputFormatContext(); err != nil {
		err = fmt.Errorf("avgo_test: getting video input format context failed: %w", err)
		return
	}

	ss := fc.Streams()
	if len(ss) < 2 {
		err = fmt.Errorf("avgo_test: invalid streams len: %d", len(ss))
		return
	}

	s1 = ss[0]
	s2 = ss[1]

	global.inputStream1 = s1
	global.inputStream2 = s2
	return
}

func TestStream(t *testing.T) {
	_, s1, s2, err := videoInputStreams()
	require.NoError(t, err)

	require.Equal(t, 0, s1.Index())
	require.Equal(t, avgo.NewRational(24, 1), s1.AvgFrameRate())
	require.Equal(t, int64(61440), s1.Duration())
	require.True(t, s1.EventFlags().Has(avgo.StreamEventFlag(2)))
	require.Equal(t, 1, s1.ID())
	require.Equal(t, "und", s1.Metadata().Get("language", nil, avgo.NewDictionaryFlags()).Value())
	require.Equal(t, int64(120), s1.NbFrames())
	require.Equal(t, avgo.NewRational(24, 1), s1.RFrameRate())
	require.Equal(t, avgo.NewRational(1, 1), s1.SampleAspectRatio())
	require.Equal(t, []byte{}, s1.SideData(avgo.PacketSideDataTypeNb))
	require.Equal(t, int64(0), s1.StartTime())
	require.Equal(t, avgo.NewRational(1, 12288), s1.TimeBase())

	require.Equal(t, 1, s2.Index())
	require.Equal(t, int64(240640), s2.Duration())
	require.Equal(t, 2, s2.ID())
	require.Equal(t, int64(235), s2.NbFrames())
	require.Equal(t, int64(0), s2.StartTime())
	require.Equal(t, avgo.NewRational(1, 48000), s2.TimeBase())

	s1.SetTimeBase(avgo.NewRational(1, 1))
	require.Equal(t, avgo.NewRational(1, 1), s1.TimeBase())
}
