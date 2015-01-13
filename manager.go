package main

//Send to subscribe to events.
//MUST handle all passed to it until channel has successfully passed to Unsubscribe
var Subscribe = make(chan chan *Event)

//Send Conn to unsubscribe.
var Unsubscribe = make(chan chan *Event)

//Add an event to event queue.
var NewEvent = make(chan *Event)

//An event to update the value of a device.
type Event struct {
	Source chan *Event
	Name   string
	Value  string
}

//Manages new connections, deleting connections, and broadcasting events
//to all of the current connections.
func Manager() {
	//Channel to get from the front of the event queue
	NextEvent := make(chan *Event)
	go func() {
		//Event queue handled as an array.  When a new event is present,
		//append it, send the first available event over NextEvent.
		//When there's no events left to send, reset the queue list to empty.
		var queue []*Event
		var next int
		for {
			if len(queue) == next {
				next = 0
				queue = queue[0:0]
				queue = append(queue, <-NewEvent)
			}
			select {
			case NextEvent <- queue[next]:
				queue[next] = nil
				next += 1
			case add := <-NewEvent:
				queue = append(queue, add)
			}
		}
	}()

	//Set of current connections.
	subscriptions := make(map[chan *Event]struct{})

	//Handle events forever
	for {
		select {
		case conn := <-Subscribe:
			connections[conn] = struct{}{}
		case conn := <-Unsubscribe:
			delete(connections, conn)
		case event := <-NextEvent:
			for subscription := range subscriptions {
				subscription <- event
			}
		}
	}
}
