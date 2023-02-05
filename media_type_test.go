package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestMediaType(t *testing.T) {
	require.Equal(t, "video", avgo.MediaTypeVideo.String())
}
