package avgo_test

import (
	"testing"

	"github.com/bubbajoe/avgo"
	"github.com/stretchr/testify/require"
)

type logItem struct {
	fmt string
	l   avgo.LogLevel
	msg string
}

func TestLog(t *testing.T) {
	var lis []logItem
	avgo.SetLogCallback(func(l avgo.LogLevel, fmt, msg, parent string) {
		lis = append(lis, logItem{
			fmt: fmt,
			l:   l,
			msg: msg,
		})
	})
	avgo.SetLogLevel(avgo.LogLevelWarning)
	avgo.Log(avgo.LogLevelDebug, "debug")
	avgo.Log(avgo.LogLevelVerbose, "verbose")
	avgo.Log(avgo.LogLevelInfo, "info")
	avgo.Log(avgo.LogLevelWarning, "warning")
	avgo.Log(avgo.LogLevelError, "error")
	avgo.Log(avgo.LogLevelFatal, "fatal")
	require.Equal(t, []logItem{
		{
			fmt: "warning",
			l:   avgo.LogLevelWarning,
			msg: "warning",
		},
		{
			fmt: "error",
			l:   avgo.LogLevelError,
			msg: "error",
		},
		{
			fmt: "fatal",
			l:   avgo.LogLevelFatal,
			msg: "fatal",
		},
	}, lis)
	avgo.ResetLogCallback()
	lis = []logItem{}
	avgo.Log(avgo.LogLevelError, "test error log\n")
	require.Equal(t, []logItem{}, lis)
}

func TestLogf(t *testing.T) {
	var lis []logItem
	avgo.SetLogCallback(func(l avgo.LogLevel, fmt, msg, parent string) {
		lis = append(lis, logItem{
			fmt: fmt,
			l:   l,
			msg: msg,
		})
	})
	avgo.SetLogLevel(avgo.LogLevelWarning)
	avgo.Logf(avgo.LogLevelDebug, "debug %s %d %.3f", "s", 1, 2.0)
	avgo.Logf(avgo.LogLevelVerbose, "verbose %s %d %.3f", "s", 1, 2.0)
	avgo.Logf(avgo.LogLevelInfo, "info %s %d %.3f", "s", 1, 2.0)
	avgo.Logf(avgo.LogLevelWarning, "warning %s %d %.3f", "s", 1, 2.0)
	avgo.Logf(avgo.LogLevelError, "error %s %d %.3f", "s", 1, 2.0)
	avgo.Logf(avgo.LogLevelFatal, "fatal %s %d %.3f", "s", 1, 2.0)
	for i, l := range []logItem{
		{
			fmt: "warning s 1 2.000",
			l:   avgo.LogLevelWarning,
			msg: "warning s 1 2.000",
		},
		{
			fmt: "error s 1 2.000",
			l:   avgo.LogLevelError,
			msg: "error s 1 2.000",
		},
		{
			fmt: "fatal s 1 2.000",
			l:   avgo.LogLevelFatal,
			msg: "fatal s 1 2.000",
		},
	} {
		require.Equal(t, l, lis[i])
	}
	avgo.ResetLogCallback()
	lis = []logItem{}
	avgo.Log(avgo.LogLevelError, "test error log\n")
	require.Equal(t, []logItem{}, lis)
}
