package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestCodecID(t *testing.T) {
	require.Equal(t, avgo.MediaTypeVideo, avgo.CodecIDH264.MediaType())
	require.Equal(t, "h264", avgo.CodecIDH264.Name())
	require.Equal(t, "h264", avgo.CodecIDH264.String())
}
