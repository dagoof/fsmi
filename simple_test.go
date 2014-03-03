package fsm

import (
	"fmt"
)

const (
	Nothing State = iota
	Ringing
	PickedUp
	Talking
	AnsweringMachine
	HungUp
)

type Phone struct {
	State
}

func (p *Phone) Current() State {
	return p.State
}

func (p *Phone) ToStates() States {
	switch p.Current() {
	case Nothing:
		return States{Ringing}
	case Ringing:
		return States{PickedUp, AnsweringMachine}
	case PickedUp:
		return States{Talking}
	case Talking:
		return States{HungUp}
	case AnsweringMachine:
		return States{PickedUp, HungUp}
	}
	return States{}
}

func (p *Phone) Set(target State) {
	p.State = target
}

func (p *Phone) TransitionTo(target State) MachineMaker {
	return Self(p, target)
}

func ExampleSimple() {
	var phone Phone

	fmt.Println(phone)

	fmt.Println(Transition(&phone, Ringing))
	fmt.Println(Transition(&phone, Ringing))
	fmt.Println(Transition(&phone, PickedUp))
	fmt.Println(Transition(&phone, Ringing))
	fmt.Println(Transition(&phone, Talking))
	fmt.Println(Transition(&phone, HungUp))

	// Output:
	// A
}
