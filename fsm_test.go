package fsmi

import (
	"testing"
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

func TestMachine(t *testing.T) {
	var phone Machine = &Phone{Nothing}
	var next State
	var err error

	next = Ringing
	phone, err = Transition(phone, next)
	if err != nil {
		t.Error(err)
	}
	if phone.Current() != next {
		t.Errorf("phone state should be %d, not %d", next, phone.Current())
	}

	next = Ringing
	phone, err = Transition(phone, next)
	if err == nil {
		t.Error(err)
	}

	next = AnsweringMachine
	phone, err = Transition(phone, next)
	if err != nil {
		t.Error(err)
	}
	if phone.Current() != next {
		t.Errorf("phone state should be %d, not %d", next, phone.Current())
	}

	next = PickedUp
	phone, err = Transition(phone, next)
	if err != nil {
		t.Error(err)
	}
	if phone.Current() != next {
		t.Errorf("phone state should be %d, not %d", next, phone.Current())
	}

	next = Talking
	phone, err = Transition(phone, next)
	if err != nil {
		t.Error(err)
	}
	if phone.Current() != next {
		t.Errorf("phone state should be %d, not %d", next, phone.Current())
	}

	next = HungUp
	phone, err = Transition(phone, next)
	if err != nil {
		t.Error(err)
	}
	if phone.Current() != next {
		t.Errorf("phone state should be %d, not %d", next, phone.Current())
	}

	next = Nothing
	phone, err = Transition(phone, next)
	if err == nil {
		t.Error(err)
	}
	if phone.Current() == next {
		t.Errorf("phone state should be not have progressed to %d", next)
	}

}

// Essentially duplicates the above test but with output for godoc. Without the
// above test, godoc writes out all the Phone implementation.
func ExampleMachine() {
	var phone Machine = &Phone{Nothing}

	fmt.Println(Transition(phone, Ringing))
	fmt.Println(Transition(phone, Ringing))
	fmt.Println(Transition(phone, AnsweringMachine))
	fmt.Println(Transition(phone, PickedUp))
	fmt.Println(Transition(phone, Talking))
	fmt.Println(Transition(phone, HungUp))
	fmt.Println(Transition(phone, Nothing))
	// Output:
	// &{1} <nil>
	// &{1} could not transition from state 1 to 1
	// &{4} <nil>
	// &{2} <nil>
	// &{3} <nil>
	// &{5} <nil>
	// &{5} could not transition from state 5 to 0
}

