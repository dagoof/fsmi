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

func (p *Phone) Current() Identifier {
	return p.ID
}

func (p *Phone) Available() []Identifier {
	switch p.Current() {
	case Nothing:
		return []Identifier{Ringing}
	case Ringing:
		return []Identifier{PickedUp, AnsweringMachine}
	case PickedUp:
		return []Identifier{Talking}
	case Talking:
		return []Identifier{HungUp}
	case AnsweringMachine:
		return []Identifier{PickedUp, HungUp}
	}
	return []Identifier{}
}

func (p *Phone) Transition(target Identifier) (Transitioner, error) {
	p.ID = ID(target.Identity())
	return p, nil
}

func TestState(t *testing.T) {
	var phone Transitioner = &Phone{Nothing}
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
	var phone Transitioner = &Phone{Nothing}

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
