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

//Analog is an analog device on the vrpn server.
type Analog struct {
	c        unsafe.Pointer
	channels int
}

//Creates a new analog device with the given device name, and the number
//of channels which the device transmits.
func (c *Connection) NewAnalog(name string, channels int) *Analog {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return &Analog{C.vrpn_Analog_Web_New(cname, c.c, C.int(channels)), channels}
}

//Updates the values on the analog device.  If there are fewer values in
//the slice than channels, only the lower channels are updated; extra
//channels are ignored.
func (a *Analog) Update(data []float64) {
	for i := 0; i < len(data) && i < a.channels; i++ {
		C.vrpn_Analog_Web_Update(a.c, C.double(data[i]), C.int(i))
	}
}

//Mainloop must be called after providing an update to the device and
//before the connection's mainloop is called.
func (a *Analog) Mainloop() {
	C.vrpn_Analog_Web_Mainloop(a.c)
}
