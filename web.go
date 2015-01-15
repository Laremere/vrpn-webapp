package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

//WebSocketHandler provides a handler for websocket connections given
//a list of devices.
type WebSocketHandler struct {
	devices []DeviceConfig
}

//HandleSocket sends devices and events over the websocket,
//and recieves the events from the websocket to broadcast.
func (wsh *WebSocketHandler) HandleSocket(ws *websocket.Conn) {
	log.Println("Websocket connected: ", ws.RemoteAddr())
	defer log.Println("Websocket disconnected", ws.RemoteAddr())

	jm := NewJSONMessenger(ws)
	for _, device := range wsh.devices {
		err := jm.Send("new_"+device.WebType(), device)
		if err != nil {
			return
		}
	}

	events := make(chan *Event)
	Subscribe <- events

	go wsh.SendEvents(events, jm)

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

//SendEvents listens to event broadcasts and sends those
//events to the web page.  It also handles closing the subscription
//when a connection is lost.
func (*WebSocketHandler) SendEvents(events chan *Event, jm *JSONMessenger) {
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

//JSONMessenger sends json objects over a writer (the websocket)
//in single calls to the write function for the whole object, so
//it can be recieved by the client as a single message.
type JSONMessenger struct {
	w io.Writer
	b bytes.Buffer
	e *json.Encoder
}

//NewJSONMessenger creates a JSONMessenger over the given writer.
func NewJSONMessenger(w io.Writer) *JSONMessenger {
	var j JSONMessenger
	j.w = w
	j.e = json.NewEncoder(&j.b)
	return &j
}

//Send sends over the writer a JSON object which has a MessageType field with the value
//given by messageType, and it serializes messageData into json in the Data field.
func (j *JSONMessenger) Send(messageType string, messageData interface{}) error {
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
