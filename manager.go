package main

//Send to subscribe to events.
//MUST handle all passed to it until channel has successfully passed to Unsubscribe
var Subscribe = make(chan chan *Event)

//Send Conn to unsubscribe.
var Unsubscribe = make(chan chan *Event)

//Add an event to event queue.
var NewEvent = make(chan *Event)

//Event to update the value of a device.
type Event struct {
	Source chan *Event
	Name   string
	Value  string
}

//Manager manages new connections, deleting connections, and broadcasting events
//to all of the current connections.
func Manager(devices []DeviceConfig) {
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
				next++
			case add := <-NewEvent:
				queue = append(queue, add)
			}
		}
	}()

	//Current state of each device
	current := make(map[string]string)
	for _, device := range devices {
		current[device.GetName()] = device.GetInitial()
	}

	//Set of current connections.
	subscriptions := make(map[chan *Event]struct{})

	//Handle events forever
	for {
		select {
		case conn := <-Subscribe:
			subscriptions[conn] = struct{}{}
			for name, val := range current {
				conn <- &Event{nil, name, val}
			}
		case conn := <-Unsubscribe:
			close(conn)
			delete(subscriptions, conn)
		case event := <-NextEvent:
			current[event.Name] = event.Value
			for subscription := range subscriptions {
				subscription <- event
			}
		}
	}
}
