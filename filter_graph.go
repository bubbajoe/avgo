package avgo

//#cgo pkg-config: libavfilter
//#include <libavfilter/avfilter.h>
import "C"
import (
	"strings"
	"unsafe"
)

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavfilter/avfilter.h#L861
type FilterGraph struct {
	c *C.struct_AVFilterGraph
}

func newFilterGraphFromC(c *C.struct_AVFilterGraph) *FilterGraph {
	if c == nil {
		return nil
	}
	return &FilterGraph{c: c}
}

func AllocFilterGraph() *FilterGraph {
	return newFilterGraphFromC(C.avfilter_graph_alloc())
}

func (g *FilterGraph) Free() {
	C.avfilter_graph_free(&g.c)
}

func (g *FilterGraph) String() string {
	return C.GoString(C.avfilter_graph_dump(g.c, nil))
}

type FilterArgs map[string]string

func (args FilterArgs) String() string {
	var ss []string
	for k, v := range args {
		ss = append(ss, k+"="+v)
	}
	return strings.Join(ss, ":")
}

func (g *FilterGraph) NewFilterContext(f *Filter, name string, args FilterArgs) (*FilterContext, error) {
	return g.NewFilterContextRawArgs(f, name, args.String())
}

func (g *FilterGraph) NewFilterContextRawArgs(f *Filter, name, args string) (*FilterContext, error) {
	ca := (*C.char)(nil)
	if len(args) > 0 {
		ca = C.CString(args)
		defer C.free(unsafe.Pointer(ca))
	}
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	fc := newFilterContext()
	err := newError(C.avfilter_graph_create_filter(&fc.c, f.c, cn, ca, nil, g.c))
	return fc, err
}

// func (g *FilterGraph) NewFilterContextByName(fc *FilterContext, name string, args FilterArgs) (*FilterContext, error) {
// 	ca := (*C.char)(nil)
// 	if args != nil && len(args) > 0 {
// 		ca = C.CString(args.String())
// 		defer C.free(unsafe.Pointer(ca))
// 	}
// 	cn := C.CString(name)
// 	defer C.free(unsafe.Pointer(cn))
// 	nfc := FindFilterByName(name)
// 	err := newError(C.avfilter_graph_create_filter(&fc.c, nfc.c, nil, ca, nil, g.c))
// 	return nfc, err
// }

func (g *FilterGraph) Parse(content string, inputs, outputs *FilterInOut) error {
	cc := C.CString(content)
	defer C.free(unsafe.Pointer(cc))
	var ic **C.struct_AVFilterInOut
	if inputs != nil {
		ic = &inputs.c
	}
	var oc **C.struct_AVFilterInOut
	if outputs != nil {
		oc = &outputs.c
	}
	return newError(C.avfilter_graph_parse_ptr(g.c, cc, ic, oc, nil))
}

func (g *FilterGraph) Configure() error {
	return newError(C.avfilter_graph_config(g.c, nil))
}

func (g *FilterGraph) SendCommand(target, cmd, args string, f FilterCommandFlags) (response string, err error) {
	targetc := C.CString(target)
	defer C.free(unsafe.Pointer(targetc))
	cmdc := C.CString(cmd)
	defer C.free(unsafe.Pointer(cmdc))
	argsc := C.CString(args)
	defer C.free(unsafe.Pointer(argsc))
	response, err = stringFromC(255, func(buf *C.char, size C.size_t) error {
		return newError(C.avfilter_graph_send_command(g.c, targetc, cmdc, argsc, buf, C.int(size), C.int(f)))
	})
	return
}
