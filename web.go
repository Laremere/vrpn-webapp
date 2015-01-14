package main

import (
	"bufio"
	"golang.org/x/net/websocket"
	"log"
)

func HandleSocket(ws *websocket.Conn) {
	log.Println("Websocket connected: ", ws.RemoteAddr())
	defer log.Println("Websocket disconnected", ws.RemoteAddr())
	events := make(chan *Event)
	Subscribe <- events

	go SendEvents(events, ws)

	r := bufio.NewReader(ws)
	for {
		name, err := r.ReadString(';')
		if err != nil {
			return
		}
		val, err := r.ReadString(';')
		if err != nil {
			return
		}
		NewEvent <- &Event{events, name[:len(name)-1], val[:len(val)-1]}
	}
}

func SendEvents(events chan *Event, ws *websocket.Conn) {
	//When send events exists, unsubscribe and drain the events
	defer func() {
		go func() {
			Unsubscribe <- events
		}()
		for range events {
		}
	}()

}
