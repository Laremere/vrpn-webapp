package main

import (
	"github.com/Laremere/vrpn-webapp/vrpn"
	"log"
	"strconv"
	"time"
)

//VrpnType is how the device should be prepresented in VRPN
type VrpnType uint8

const (
	//VrpnButton represents a vrpn true or false value
	VrpnButton = VrpnType(iota)
	//VrpnAnalog represents a vrpn float value
	VrpnAnalog
)

//VrpnServe handles starting a vrpn server and broadcasting events over it.
func VrpnServe(port uint16, devices []DeviceConfig) {
	conn := vrpn.NewConnection(int(port))

	vrpnButtonDevices := make(map[string]*vrpn.Button)
	vrpnAnalogDevices := make(map[string]*vrpn.Analog)
	for _, device := range devices {
		name := device.GetName()
		if device.VrpnType() == VrpnButton {
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
			} else if analog, ok := vrpnAnalogDevices[event.Name]; ok {
				val, err := strconv.ParseFloat(event.Value, 64)
				if err != nil {
					log.Println("Error parsing analog value,", err)
					continue
				}
				analog.Update([]float64{val})
			} else {
				log.Println("Unkown device identity,", event.Name)
			}
		case <-ticker:
		}
		for _, button := range vrpnButtonDevices {
			button.Mainloop()
		}
		for _, analog := range vrpnAnalogDevices {
			analog.Mainloop()
		}
		conn.Mainloop()
	}
}
