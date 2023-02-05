package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestChannelLayout(t *testing.T) {
	require.Equal(t, 2, avgo.ChannelLayoutStereo.NbChannels())
	require.Equal(t, "stereo", avgo.ChannelLayoutStereo.String())
	require.Equal(t, "1 channels (FL+FR)", avgo.ChannelLayoutStereo.StringWithNbChannels(1))
}
