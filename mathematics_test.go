package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestMathematics(t *testing.T) {
	require.Equal(t, int64(1000), avgo.RescaleQ(100, avgo.NewRational(1, 100), avgo.NewRational(1, 1000)))
	require.Equal(t, int64(0), avgo.RescaleQRnd(1, avgo.NewRational(1, 100), avgo.NewRational(1, 10), avgo.RoundingDown))
	require.Equal(t, int64(1), avgo.RescaleQRnd(1, avgo.NewRational(1, 100), avgo.NewRational(1, 10), avgo.RoundingUp))
}
