/*
Package vrpn provides a wrapper around a simple vrpn server.

All types in this package wrap around the required vrpn_wrapper.dll.
It is important to know that objects are created in the C++ side, and
will not be automatically garbage collected, even if the go objects are.

*/
package vrpn

/*
#cgo CPPFLAGS: -I vrpn_wrapper/vrpn_wrapper
#cgo LDFLAGS: vrpn_wrapper/Release/vrpn_wrapper.dll
#include "vrpn_Connection_Wrapper.h"

*/
import "C"
import (
	"unsafe"
)

//Connection represents a vrpn server.
type Connection struct {
	c unsafe.Pointer
}

//Creates a new vrpn server at the given port number.
func NewConnection(port int) *Connection {
	return &Connection{C.vrpn_Connection_New(C.int(port))}
}

//Mainloop must be called each loop after each device's main loop in order to
//ensure that vrpn does all it's nessisary transmissions and bookkeeping.
func (c *Connection) Mainloop() {
	C.vrpn_Connection_Mainloop(c.c)
}
