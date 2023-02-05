package avgo

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

type StreamEventFlag int

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavformat/avformat.h#L1070
const (
	StreamEventFlagMetadataUpdated = StreamEventFlag(C.AVSTREAM_EVENT_FLAG_METADATA_UPDATED)
)
