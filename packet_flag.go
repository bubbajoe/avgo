package avgo

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

type PacketFlag int

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavcodec/packet.h#L428
const (
	PacketFlagCorrupt = PacketFlag(C.AV_PKT_FLAG_CORRUPT)
	PacketFlagDiscard = PacketFlag(C.AV_PKT_FLAG_DISCARD)
	PacketFlagKey     = PacketFlag(C.AV_PKT_FLAG_KEY)
)
