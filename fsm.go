package fsmi

import (
	"fmt"
)

// State is a type meant to be declared in an iota and used as an enum to mark
// transitions between machines.
type State int64

// States is a slice of states
type States []State

// Machine is an interface that describes one part of a 
type Node interface {
	Current() State
	ToStates() States
	Transition(State) Machine
}

type MachineMaker func(Machine) Machine

func Transition(m Machine, target State) (Machine, error) {
	if CanTransition(m, target) {
		return m.Transition(target), nil
	}
	return m, TransitionError{m.Current(), target}
}

func CanTransition(m Machine, target State) bool {
	for _, available := range m.ToStates() {
		if target == available {
			return true
		}
	}

	return false
}

type TransitionError struct {
	From, To State
}

func (t TransitionError) Error() string {
	return fmt.Sprintf("could not transition from state %d to %d", t.From, t.To)
}
