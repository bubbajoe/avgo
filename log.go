package avgo

//#cgo pkg-config: libavutil
//#include <stdlib.h>
//#include <libavutil/log.h>
/*
#include <stdio.h>

extern void goavgoLogCallback(int level, char* fmt, char* msg, char* parent);

static inline void avgoLogCallback(void *avcl, int level, const char *fmt, va_list vl)
{
	if (level > av_log_get_level()) return;
	AVClass* avc = avcl ? *(AVClass **) avcl : NULL;
	char parent[1024];
	if (avc) {
		sprintf(parent, "%p", avcl);
	}
	char msg[1024];
	vsprintf(msg, fmt, vl);
	goavgoLogCallback(level, (char*)(fmt), msg, parent);
}
static inline void avgoSetLogCallback()
{
	av_log_set_callback(avgoLogCallback);
}
static inline void avgoResetLogCallback()
{
	av_log_set_callback(av_log_default_callback);
}
static inline void avgoLog(int level, const char *fmt)
{
	av_log(NULL, level, fmt, NULL);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type LogLevel int

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavutil/log.h#L162
const (
	LogLevelQuiet   = LogLevel(C.AV_LOG_QUIET)
	LogLevelPanic   = LogLevel(C.AV_LOG_PANIC)
	LogLevelFatal   = LogLevel(C.AV_LOG_FATAL)
	LogLevelError   = LogLevel(C.AV_LOG_ERROR)
	LogLevelWarning = LogLevel(C.AV_LOG_WARNING)
	LogLevelInfo    = LogLevel(C.AV_LOG_INFO)
	LogLevelVerbose = LogLevel(C.AV_LOG_VERBOSE)
	LogLevelDebug   = LogLevel(C.AV_LOG_DEBUG)
)

func SetLogLevel(l LogLevel) {
	C.av_log_set_level(C.int(l))
}

func (l LogLevel) String() string {
	switch l {
	case LogLevelQuiet:
		return "quiet"
	case LogLevelPanic:
		return "panic"
	case LogLevelFatal:
		return "fatal"
	case LogLevelError:
		return "error"
	case LogLevelWarning:
		return "warning"
	case LogLevelInfo:
		return "info"
	case LogLevelVerbose:
		return "verbose"
	case LogLevelDebug:
		return "debug"
	}
	return ""
}

type LogCallback func(l LogLevel, fmt, msg, parent string)

var logCallback LogCallback

func SetLogCallback(c LogCallback) {
	logCallback = c
	C.avgoSetLogCallback()
}

//export goavgoLogCallback
func goavgoLogCallback(level C.int, fmt, msg, parent *C.char) {
	if logCallback == nil {
		return
	}
	logCallback(LogLevel(level), C.GoString(fmt), C.GoString(msg), C.GoString(parent))
}

func ResetLogCallback() {
	C.avgoResetLogCallback()
}

func Log(l LogLevel, msg string) {
	msgc := C.CString(msg)
	defer C.free(unsafe.Pointer(msgc))
	C.avgoLog(C.int(l), msgc)
}

func Logf(l LogLevel, msg string, args ...interface{}) {
	msgc := C.CString(fmt.Sprintf(msg, args...))
	defer C.free(unsafe.Pointer(msgc))
	C.avgoLog(C.int(l), msgc)
}
