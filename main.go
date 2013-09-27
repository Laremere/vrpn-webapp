package main

import (
	"github.com/Laremere/vrpn-webapp/vrpn"
	"log"
	"runtime"
)

var devices map[string]Device

func main() {
	runtime.GOMAXPROCS(2)
	log.Println("Starting...")

	conn := vrpn.NewConnection(3883)
	devices = make(map[string]Device)
	devices["Analog0"] = NewAnalogDevice(conn, "Analog0", 1)

	button := conn.NewButton("Button0", 1)

	go StartHttp(80)

	for {
		for _, d := range devices {
			d.Mainloop()
		}

		button.Mainloop()
		conn.Mainloop()
	}
}
