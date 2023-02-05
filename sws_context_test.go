package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
)

// TestNewSwsContext
func TestNewSwsContext(t *testing.T) {
	// Init
	var (
		s = avgo.NewSwsContext(
			100, 100, avgo.PixelFormatYuv420P,
			100, 100, avgo.PixelFormatYuv420P,
			0, nil, nil,
		)
	)
	// Free
	s.Free()
}

// TestInitContext

// TestCachedContext

// TestScale

// TestScaleFrames

// TestScaleDstFrame

// TestScaleSrcFrame
