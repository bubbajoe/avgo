package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestPictureType(t *testing.T) {
	require.Equal(t, "I", avgo.PictureTypeI.String())
}
