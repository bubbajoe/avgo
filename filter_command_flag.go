package avgo

//#cgo pkg-config: libavfilter
//#include <libavfilter/avfilter.h>
import "C"

type FilterCommandFlag int

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavfilter/avfilter.h#L739
const (
	FilterCommandFlagOne  = FilterCommandFlag(C.AVFILTER_CMD_FLAG_ONE)
	FilterCommandFlagFast = FilterCommandFlag(C.AVFILTER_CMD_FLAG_FAST)
)
