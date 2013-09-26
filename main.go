package main

import (
	"github.com/Laremere/vrpn-webapp/vrpn"
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	log.Println("Starting...")

	go StartHttp(80)

	conn := vrpn.NewConnection(3883)
	button := conn.NewButton("Button0", 2)

	buttonChan := make(chan []bool)
	go func() {
		active := make([]bool, 2)
		for {
			i := rand.Int() % 2
			active[i] = !active[i]
			buttonChan <- active
			time.Sleep(time.Second * 1)
		}
	}()

	analog := conn.NewAnalog("Analog0", 2)
	analogChan := make(chan []float64)
	go func() {
		var rads float64 = 0.0
		for {
			rads += 0.1
			point := make([]float64, 2)
			point[0] = math.Cos(rads)
			point[1] = math.Sin(rads)
			analogChan <- point
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		select {
		case active := <-buttonChan:
			button.Update(active)
		case point := <-analogChan:
			analog.Update(point)
		}

		analog.Mainloop()
		button.Mainloop()
		conn.Mainloop()
	}
}
