package main

import (
	"github.com/Laremere/vrpn-webapp/vrpn"
	"log"
	"strconv"
	"time"
)

type vrpnType uint8

const (
	vrpnButton = vrpnType(iota)
	vrpnAnalog
)

func vrpnServe(port uint16, devices []DeviceConfig) {
	conn := vrpn.NewConnection(int(port))

	vrpnButtonDevices := make(map[string]*vrpn.Button)
	vrpnAnalogDevices := make(map[string]*vrpn.Analog)
	for _, device := range devices {
		name := device.GetName()
		if device.VrpnType() == vrpnButton {
			vrpnButtonDevices[name] = conn.NewButton(name, 1)
		} else {
			vrpnAnalogDevices[name] = conn.NewAnalog(name, 1)
		}
	}

	events := make(chan *Event)
	Subscribe <- events

	ticker := time.Tick(time.Second)
	for {
		select {
		case event := <-events:
			if button, ok := vrpnButtonDevices[event.Name]; ok {
				val, err := strconv.ParseBool(event.Value)
				if err != nil {
					log.Println("Error parsing button value,", err)
					continue
				}
				button.Update([]bool{val})
				button.Mainloop()
			} else {
				analog := vrpnAnalogDevices[event.Name]
				val, err := strconv.ParseFloat(event.Value, 64)
				if err != nil {
					log.Println("Error parsing analog value,", err)
					continue
				}
				analog.Update([]float64{val})
				analog.Mainloop()
			}
		case <-ticker:
		}
		conn.Mainloop()
	}
}
