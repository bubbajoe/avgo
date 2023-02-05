package avgo_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

type testError struct{}

func (err testError) Error() string { return "" }

func TestError(t *testing.T) {
	require.Equal(t, "Decoder not found", avgo.ErrDecoderNotFound.Error())
	err1 := fmt.Errorf("test 1: %w", avgo.ErrDecoderNotFound)
	require.True(t, errors.Is(err1, avgo.ErrDecoderNotFound))
	require.False(t, errors.Is(err1, testError{}))
	err2 := fmt.Errorf("test 2: %w", avgo.ErrDemuxerNotFound)
	require.False(t, errors.Is(err2, avgo.ErrDecoderNotFound))
}
