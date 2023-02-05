package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	f := avgo.FindFilterByName("format")
	require.NotNil(t, f)
	require.Equal(t, "format", f.Name())
	require.Equal(t, "format", f.String())
}
