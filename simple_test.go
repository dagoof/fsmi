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

func (p *Phone) Transition(target State) Machine {
	p.State = target
	return p
}

func ExampleMachine() {
	var phone Machine = &Phone{Nothing}
	var err error

	phone, err = Transition(phone, Ringing)
	fmt.Println(phone.Current(), err)

	phone, err = Transition(phone, Ringing)
	fmt.Println(phone.Current(), err)

	phone, err = Transition(phone, AnsweringMachine)
	fmt.Println(phone.Current(), err)

	phone, err = Transition(phone, PickedUp)
	fmt.Println(phone.Current(), err)

	phone, err = Transition(phone, Talking)
	fmt.Println(phone.Current(), err)

	phone, err = Transition(phone, HungUp)
	fmt.Println(phone.Current(), err)

	phone, err = Transition(phone, Nothing)
	fmt.Println(phone.Current(), err)

	// Output:
	// 1 <nil>
	// 1 could not transition from state 1 to 1
	// 4 <nil>
	// 2 <nil>
	// 3 <nil>
	// 5 <nil>
	// 5 could not transition from state 5 to 0
}
