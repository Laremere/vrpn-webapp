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

type Connection struct {
	c unsafe.Pointer
}

func NewConnection(port int) Connection {
	return Connection{C.vrpn_Connection_New(C.int(port))}
}

func (c Connection) Mainloop() {
	C.vrpn_Connection_Mainloop(c.c)
}
