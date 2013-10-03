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

//Button is an button device on the vrpn server.
type Button struct {
	c        unsafe.Pointer
	channels int
}

//Creates a new button device with the given device name, and the number
//of channels which the device transmits.
func (c *Connection) NewButton(name string, channels int) *Button {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return &Button{C.vrpn_Button_Web_New(cname, c.c, C.int(channels)), channels}
}

//Updates the values on the button device.  If there are fewer values in
//the slice than channels, only the lower channels are updated; extra
//channels are ignored.
func (b *Button) Update(data []bool) {
	for i := 0; i < len(data) && i < b.channels; i++ {
		var active C.char = 0
		if data[i] {
			active = 1
		}
		C.vrpn_Button_Web_Update(b.c, active, C.int(i))
	}
}

//Mainloop must be called after providing an update to the device and
//before the connection's mainloop is called.
func (b *Button) Mainloop() {
	C.vrpn_Button_Web_Mainloop(b.c)
}
