package main

import (
	"github.com/Laremere/vrpn-webapp/vrpn"
	"log"
	"strconv"
)

type Device interface {
	Update([]string)
	Mainloop()
}

type AnalogDevice struct {
	a      *vrpn.Analog
	loop   chan chan bool
	update chan []string
}

func NewAnalogDevice(conn *vrpn.Connection, name string, channels int) *AnalogDevice {
	d := new(AnalogDevice)
	d.a = conn.NewAnalog(name, channels)
	d.loop = make(chan chan bool)
	d.update = make(chan []string)
	go d.main(channels)
	return d
}

func (d *AnalogDevice) main(channels int) {
	data := make([]float64, channels)
	for {
		select {
		case message := <-d.update:
			for i := 0; i+1 < len(message); i += 2 {
				index, err := strconv.ParseInt(message[i], 10, 64)
				if err != nil {
					log.Println("Invalid analog index: ", message[i])
					continue
				}

				val, err := strconv.ParseFloat(message[i+1], 64)
				if err != nil {
					log.Println("Invalid analog value: ", message[i+1])
					continue
				}

				data[index] = val
			}
		case wait := <-d.loop:
			d.a.Update(data)
			d.a.Mainloop()
			wait <- true
		}
	}
}

func (d *AnalogDevice) Mainloop() {
	wait := make(chan bool)
	d.loop <- wait
	<-wait
}

func (d *AnalogDevice) Update(message []string) {
	d.update <- message
}
