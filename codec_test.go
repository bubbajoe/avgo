package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestCodec(t *testing.T) {
	c := avgo.FindDecoder(avgo.CodecIDMp3)
	require.NotNil(t, c)
	require.Nil(t, c.ChannelLayouts())
	require.True(t, c.IsDecoder())
	require.False(t, c.IsEncoder())
	require.Nil(t, c.PixelFormats())
	require.Equal(t, []avgo.SampleFormat{avgo.SampleFormatFltp, avgo.SampleFormatFlt}, c.SampleFormats())
	require.Equal(t, "mp3float", c.Name())
	require.Equal(t, "mp3float", c.String())

	c = avgo.FindDecoderByName("aac")
	require.NotNil(t, c)
	require.Equal(t, []avgo.ChannelLayout{
		avgo.ChannelLayoutMono,
		avgo.ChannelLayoutStereo,
		avgo.ChannelLayoutSurround,
		avgo.ChannelLayout4Point0,
		avgo.ChannelLayout5Point0Back,
		avgo.ChannelLayout5Point1Back,
		avgo.ChannelLayout7Point1WideBack,
	}, c.ChannelLayouts())
	require.True(t, c.IsDecoder())
	require.False(t, c.IsEncoder())
	require.Equal(t, []avgo.SampleFormat{avgo.SampleFormatFltp}, c.SampleFormats())
	require.Equal(t, "aac", c.Name())
	require.Equal(t, "aac", c.String())

	c = avgo.FindEncoder(avgo.CodecIDH264)
	require.NotNil(t, c)
	require.False(t, c.IsDecoder())
	require.True(t, c.IsEncoder())
	require.Equal(t, []avgo.PixelFormat{
		avgo.PixelFormatVideotoolbox,
		avgo.PixelFormatNv12,
		avgo.PixelFormatYuv420P,
	}, c.PixelFormats())
	require.Nil(t, c.SampleFormats())
	require.Equal(t, "h264_videotoolbox", c.Name())
	require.Equal(t, "h264_videotoolbox", c.String())

	c = avgo.FindEncoderByName("mjpeg")
	require.NotNil(t, c)
	require.False(t, c.IsDecoder())
	require.True(t, c.IsEncoder())
	require.Equal(t, []avgo.PixelFormat{
		avgo.PixelFormatYuvj420P,
		avgo.PixelFormatYuvj422P,
		avgo.PixelFormatYuvj444P,
	}, c.PixelFormats())
	require.Equal(t, "mjpeg", c.Name())
	require.Equal(t, "mjpeg", c.String())

	c = avgo.FindDecoderByName("invalid")
	require.Nil(t, c)
}
