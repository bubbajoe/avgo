package avgo

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

type FormatContextCtxFlag int

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavformat/avformat.h#L1153
const (
	FormatContextCtxFlagNoHeader   = FormatContextCtxFlag(C.AVFMTCTX_NOHEADER)
	FormatContextCtxFlagUnseekable = FormatContextCtxFlag(C.AVFMTCTX_UNSEEKABLE)
)
