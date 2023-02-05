package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestSampleFormat(t *testing.T) {
	require.Equal(t, "s16", avgo.SampleFormatS16.String())
}
