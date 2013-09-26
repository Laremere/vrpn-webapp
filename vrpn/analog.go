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
	c        unsafe.Pointer
	channels int
}

func (c Connection) NewAnalog(name string, channels int) Analog {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Analog{C.vrpn_Analog_Web_New(cname, c.c, C.int(channels)), channels}
}

func (a Analog) Update(data []float64) {
	for i := 0; i < len(data) && i < a.channels; i++ {
		C.vrpn_Analog_Web_Update(a.c, C.double(data[i]), C.int(i))
	}
}

func (a Analog) Mainloop() {
	C.vrpn_Analog_Web_Mainloop(a.c)
}
