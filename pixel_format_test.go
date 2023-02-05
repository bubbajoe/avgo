package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestPixelFormat(t *testing.T) {
	p := avgo.FindPixelFormatByName("yuv420p")
	require.Equal(t, avgo.PixelFormatYuv420P, p)
	require.Equal(t, "yuv420p", p.String())
}
