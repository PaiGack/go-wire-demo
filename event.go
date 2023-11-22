package main

import "fmt"

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}
