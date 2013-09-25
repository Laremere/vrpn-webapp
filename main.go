package main

import (
	"github.com/Laremere/vrpn-webapp/vrpn"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")
	conn := vrpn.NewConnection()
	button := conn.NewButton("Button0", 1)
	log.Println(conn)

	buttonChan := make(chan bool)
	go func() {
		var active bool
		for {
			active = !active
			buttonChan <- active
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		select {
		case active := <-buttonChan:
			button.Update([]bool{active})
		}

		button.Mainloop()
		conn.Mainloop()
	}
}
