package vrpn

/*
#cgo CPPFLAGS: -I vrpn_wrapper/vrpn_wrapper
#cgo LDFLAGS: ../vrpn_wrapper.dll
#include "vrpn_Analog_Web_Wrapper.h"
#include "stdlib.h"

*/
import "C"
import (
	"unsafe"
)

type Analog struct {
	c unsafe.Pointer
}

func (c Connection) NewAnalog(name string, channels int) Analog {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Analog{C.vrpn_Analog_Web_New(cname, c.c, C.int(channels))}
}

func (b Analog) Update(data []float64) {
	for i := 0; i < len(data); i++ {
		C.vrpn_Analog_Web_Update(b.c, C.double(data[i]), C.int(i))
	}
}

func (b Analog) Mainloop() {
	C.vrpn_Analog_Web_Mainloop(b.c)
}
