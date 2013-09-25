package vrpn

/*
#cgo CPPFLAGS: -I vrpn_wrapper/vrpn_wrapper
#cgo LDFLAGS: ../vrpn_wrapper.dll
#include "vrpn_Button_Web_Wrapper.h"
#include "stdlib.h"

*/
import "C"
import (
	"unsafe"
)

type Button struct {
	c unsafe.Pointer
}

func (c Connection) NewButton(name string, channels int) Button {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Button{C.vrpn_Button_Web_New(cname, c.c, C.int(channels))}
}

func (b Button) Update(data []bool) {
	for i := 0; i < len(data); i++ {
		var active C.char = 0
		if data[i] {
			active = 1
		}
		C.vrpn_Button_Web_Update(b.c, active, C.int(i))
	}
}

func (b Button) Mainloop() {
	C.vrpn_Button_Web_Mainloop(b.c)
}
