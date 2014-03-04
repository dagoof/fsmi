package fsmi

import (
	"fmt"
	"testing"
)

const (
	Nothing ID = iota
	Ringing
	PickedUp
	Talking
	AnsweringMachine
	HungUp
)

type Phone struct {
	ID
}

func (p *Phone) Current() ID {
	return p.ID
}

func (p *Phone) Available() IDs {
	switch p.Current() {
	case Nothing:
		return IDs{Ringing}
	case Ringing:
		return IDs{PickedUp, AnsweringMachine}
	case PickedUp:
		return IDs{Talking}
	case Talking:
		return IDs{HungUp}
	case AnsweringMachine:
		return IDs{PickedUp, HungUp}
	}
	return IDs{}
}

func (p *Phone) Transition(target ID) State {
	p.ID = target
	return p
}

func TestState(t *testing.T) {
	var phone State = &Phone{Nothing}
	var next ID
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

// See fsm_test.go for the Available method on Phone which defines the available
// transitions for any current state of Phone.
func ExampleState() {
	var phone State = &Phone{Nothing}

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
