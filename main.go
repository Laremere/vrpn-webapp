package main

import (
	"github.com/Laremere/vrpn-webapp/vrpn"
	"log"
	"runtime"
	"time"
)

var devices map[string]*Device

func main() {
	runtime.GOMAXPROCS(2)
	log.Println("Starting...")
	configs := ReadConfig()
	log.Println("vrpn Port:", configs.VrpnPort)
	log.Println("http Port:", configs.HttpPort)

	conn := vrpn.NewConnection(configs.VrpnPort)
	devices = make(map[string]*Device)
	for _, val := range configs.Devices {
		devices[val.Name] = CreateDevice(conn, val)
	}

	go StartHttp(configs.HttpPort)

	for {
		for _, d := range devices {
			d.Mainloop()
		}
		conn.Mainloop()
		time.Sleep(time.Second / 30)
	}
}
