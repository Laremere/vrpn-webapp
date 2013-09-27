package main

import (
	"github.com/Laremere/vrpn-webapp/vrpn"
	"log"
	"strconv"
)

type deviceInterpreter interface {
	Update([]string)
	Mainloop()
}

type Device struct {
	loop   chan chan bool
	update chan []string
	di     deviceInterpreter
}

func newDevice(conn *vrpn.Connection, di deviceInterpreter) *Device {
	d := new(Device)
	d.loop = make(chan chan bool)
	d.update = make(chan []string)
	d.di = di
	go d.main()
	return d
}

func (d *Device) main() {
	for {
		select {
		case message := <-d.update:
			d.di.Update(message)
		case wait := <-d.loop:
			d.di.Mainloop()
			wait <- true
		}
	}
}

func (d *Device) Mainloop() {
	wait := make(chan bool)
	d.loop <- wait
	<-wait
}

func (d *Device) Update(message []string) {
	d.update <- message
}

type analogDevice struct {
	a    *vrpn.Analog
	data []float64
}

func NewAnalogDevice(conn *vrpn.Connection, name string, channels int) *Device {
	ad := new(analogDevice)
	ad.a = conn.NewAnalog(name, channels)
	ad.data = make([]float64, channels)
	return newDevice(conn, ad)
}

func (ad *analogDevice) Update(message []string) {
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

		ad.data[index] = val
	}
}

func (ad *analogDevice) Mainloop() {
	ad.a.Update(ad.data)
	ad.a.Mainloop()
}

type buttonDevice struct {
	b    *vrpn.Button
	data []bool
}

func NewButtonDevice(conn *vrpn.Connection, name string, channels int) *Device {
	bd := new(buttonDevice)
	bd.b = conn.NewButton(name, channels)
	bd.data = make([]bool, channels)
	return newDevice(conn, bd)
}

func (bd *buttonDevice) Update(message []string) {
	for i := 0; i+1 < len(message); i += 2 {
		index, err := strconv.ParseInt(message[i], 10, 64)
		if err != nil {
			log.Println("Invalid button index: ", message[i])
			continue
		}

		val, err := strconv.ParseBool(message[i+1])
		if err != nil {
			log.Println("Invalid button value: ", message[i+1])
			continue
		}

		bd.data[index] = val
	}
}

func (bd *buttonDevice) Mainloop() {
	bd.b.Update(bd.data)
	bd.b.Mainloop()
}
