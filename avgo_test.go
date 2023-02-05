package avgo_test

import (
	"os"
	"testing"

	"github.com/bubbajoe/avgo"
)

var global = struct {
	closer             *avgo.Closer
	frame              *avgo.Frame
	inputFormatContext *avgo.FormatContext
	inputStream1       *avgo.Stream
	inputStream2       *avgo.Stream
	pkt                *avgo.Packet
}{
	closer: avgo.NewCloser(),
}

func TestMain(m *testing.M) {
	// Run
	m.Run()

	// Make sure to close closer
	global.closer.Close()

	// Exit
	os.Exit(0)
}
