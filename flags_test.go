// Code generated by avgo. DO NOT EDIT.
package avgo_test
import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

func TestBuffersinkFlags(t *testing.T) {
	fs := avgo.NewBuffersinkFlags(avgo.BuffersinkFlag(1))
	require.True(t, fs.Has(avgo.BuffersinkFlag(1)))
	fs = fs.Add(avgo.BuffersinkFlag(2))
	require.True(t, fs.Has(avgo.BuffersinkFlag(2)))
	fs = fs.Del(avgo.BuffersinkFlag(2))
	require.False(t, fs.Has(avgo.BuffersinkFlag(2)))
}

func TestBuffersrcFlags(t *testing.T) {
	fs := avgo.NewBuffersrcFlags(avgo.BuffersrcFlag(1))
	require.True(t, fs.Has(avgo.BuffersrcFlag(1)))
	fs = fs.Add(avgo.BuffersrcFlag(2))
	require.True(t, fs.Has(avgo.BuffersrcFlag(2)))
	fs = fs.Del(avgo.BuffersrcFlag(2))
	require.False(t, fs.Has(avgo.BuffersrcFlag(2)))
}

func TestCodecContextFlags(t *testing.T) {
	fs := avgo.NewCodecContextFlags(avgo.CodecContextFlag(1))
	require.True(t, fs.Has(avgo.CodecContextFlag(1)))
	fs = fs.Add(avgo.CodecContextFlag(2))
	require.True(t, fs.Has(avgo.CodecContextFlag(2)))
	fs = fs.Del(avgo.CodecContextFlag(2))
	require.False(t, fs.Has(avgo.CodecContextFlag(2)))
}

func TestCodecContextFlags2(t *testing.T) {
	fs := avgo.NewCodecContextFlags2(avgo.CodecContextFlag2(1))
	require.True(t, fs.Has(avgo.CodecContextFlag2(1)))
	fs = fs.Add(avgo.CodecContextFlag2(2))
	require.True(t, fs.Has(avgo.CodecContextFlag2(2)))
	fs = fs.Del(avgo.CodecContextFlag2(2))
	require.False(t, fs.Has(avgo.CodecContextFlag2(2)))
}

func TestDictionaryFlags(t *testing.T) {
	fs := avgo.NewDictionaryFlags(avgo.DictionaryFlag(1))
	require.True(t, fs.Has(avgo.DictionaryFlag(1)))
	fs = fs.Add(avgo.DictionaryFlag(2))
	require.True(t, fs.Has(avgo.DictionaryFlag(2)))
	fs = fs.Del(avgo.DictionaryFlag(2))
	require.False(t, fs.Has(avgo.DictionaryFlag(2)))
}

func TestFilterCommandFlags(t *testing.T) {
	fs := avgo.NewFilterCommandFlags(avgo.FilterCommandFlag(1))
	require.True(t, fs.Has(avgo.FilterCommandFlag(1)))
	fs = fs.Add(avgo.FilterCommandFlag(2))
	require.True(t, fs.Has(avgo.FilterCommandFlag(2)))
	fs = fs.Del(avgo.FilterCommandFlag(2))
	require.False(t, fs.Has(avgo.FilterCommandFlag(2)))
}

func TestFormatContextCtxFlags(t *testing.T) {
	fs := avgo.NewFormatContextCtxFlags(avgo.FormatContextCtxFlag(1))
	require.True(t, fs.Has(avgo.FormatContextCtxFlag(1)))
	fs = fs.Add(avgo.FormatContextCtxFlag(2))
	require.True(t, fs.Has(avgo.FormatContextCtxFlag(2)))
	fs = fs.Del(avgo.FormatContextCtxFlag(2))
	require.False(t, fs.Has(avgo.FormatContextCtxFlag(2)))
}

func TestFormatContextFlags(t *testing.T) {
	fs := avgo.NewFormatContextFlags(avgo.FormatContextFlag(1))
	require.True(t, fs.Has(avgo.FormatContextFlag(1)))
	fs = fs.Add(avgo.FormatContextFlag(2))
	require.True(t, fs.Has(avgo.FormatContextFlag(2)))
	fs = fs.Del(avgo.FormatContextFlag(2))
	require.False(t, fs.Has(avgo.FormatContextFlag(2)))
}

func TestFormatEventFlags(t *testing.T) {
	fs := avgo.NewFormatEventFlags(avgo.FormatEventFlag(1))
	require.True(t, fs.Has(avgo.FormatEventFlag(1)))
	fs = fs.Add(avgo.FormatEventFlag(2))
	require.True(t, fs.Has(avgo.FormatEventFlag(2)))
	fs = fs.Del(avgo.FormatEventFlag(2))
	require.False(t, fs.Has(avgo.FormatEventFlag(2)))
}

func TestIOContextFlags(t *testing.T) {
	fs := avgo.NewIOContextFlags(avgo.IOContextFlag(1))
	require.True(t, fs.Has(avgo.IOContextFlag(1)))
	fs = fs.Add(avgo.IOContextFlag(2))
	require.True(t, fs.Has(avgo.IOContextFlag(2)))
	fs = fs.Del(avgo.IOContextFlag(2))
	require.False(t, fs.Has(avgo.IOContextFlag(2)))
}

func TestIOFormatFlags(t *testing.T) {
	fs := avgo.NewIOFormatFlags(avgo.IOFormatFlag(1))
	require.True(t, fs.Has(avgo.IOFormatFlag(1)))
	fs = fs.Add(avgo.IOFormatFlag(2))
	require.True(t, fs.Has(avgo.IOFormatFlag(2)))
	fs = fs.Del(avgo.IOFormatFlag(2))
	require.False(t, fs.Has(avgo.IOFormatFlag(2)))
}

func TestPacketFlags(t *testing.T) {
	fs := avgo.NewPacketFlags(avgo.PacketFlag(1))
	require.True(t, fs.Has(avgo.PacketFlag(1)))
	fs = fs.Add(avgo.PacketFlag(2))
	require.True(t, fs.Has(avgo.PacketFlag(2)))
	fs = fs.Del(avgo.PacketFlag(2))
	require.False(t, fs.Has(avgo.PacketFlag(2)))
}

func TestSeekFlags(t *testing.T) {
	fs := avgo.NewSeekFlags(avgo.SeekFlag(1))
	require.True(t, fs.Has(avgo.SeekFlag(1)))
	fs = fs.Add(avgo.SeekFlag(2))
	require.True(t, fs.Has(avgo.SeekFlag(2)))
	fs = fs.Del(avgo.SeekFlag(2))
	require.False(t, fs.Has(avgo.SeekFlag(2)))
}

func TestStreamEventFlags(t *testing.T) {
	fs := avgo.NewStreamEventFlags(avgo.StreamEventFlag(1))
	require.True(t, fs.Has(avgo.StreamEventFlag(1)))
	fs = fs.Add(avgo.StreamEventFlag(2))
	require.True(t, fs.Has(avgo.StreamEventFlag(2)))
	fs = fs.Del(avgo.StreamEventFlag(2))
	require.False(t, fs.Has(avgo.StreamEventFlag(2)))
}
