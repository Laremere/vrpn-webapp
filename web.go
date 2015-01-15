package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

type WebSocketHandler struct {
	devices []DeviceConfig
}

func (wsh *WebSocketHandler) HandleSocket(ws *websocket.Conn) {
	log.Println("Websocket connected: ", ws.RemoteAddr())
	defer log.Println("Websocket disconnected", ws.RemoteAddr())

	jm := NewJsonMessenger(ws)
	for _, device := range wsh.devices {
		err := jm.Send("new_"+device.WebType(), device)
		if err != nil {
			return
		}
	}

	events := make(chan *Event)
	Subscribe <- events

	go SendEvents(events, jm)

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

func SendEvents(events chan *Event, jm *JsonMessenger) {
	//When send events exists, unsubscribe and drain the events
	defer func() {
		go func() {
			Unsubscribe <- events
		}()
		for range events {
		}
	}()

	type eventMessage struct {
		Device      string
		Value       string
		SelfSourced bool
	}

	for event := range events {
		var em eventMessage
		em.Device = event.Name
		em.Value = event.Value
		em.SelfSourced = events == event.Source
		err := jm.Send("event", &em)
		if err != nil {
			return
		}
	}

}

type JsonMessenger struct {
	w io.Writer
	b bytes.Buffer
	e *json.Encoder
}

func NewJsonMessenger(w io.Writer) *JsonMessenger {
	var j JsonMessenger
	j.w = w
	j.e = json.NewEncoder(&j.b)
	return &j
}

func (j *JsonMessenger) Send(messageType string, messageData interface{}) error {
	type Message struct {
		MessageType string
		Data        interface{}
	}

	defer j.b.Reset()
	payload := Message{messageType, messageData}
	err := j.e.Encode(&payload)
	if err != nil {
		return err
	}
	_, err = j.w.Write(j.b.Bytes())
	return err
}
