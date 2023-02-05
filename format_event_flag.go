package avgo

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

type FormatEventFlag int

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavformat/avformat.h#L1519
const (
	FormatEventFlagMetadataUpdated = FormatEventFlag(C.AVFMT_EVENT_FLAG_METADATA_UPDATED)
)
