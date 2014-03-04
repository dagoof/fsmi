// Package fsmi exposes a method of working with FSM's through interfaces rather
// than concrete state structs. By satisfying the State interface and using the
// Transition function, it is possible to have arbitrary types act behave as FSM
// states without embedding State structs everywhere and cluttering your types.
package fsmi

import (
	"fmt"
)

// ID is a type meant to be declared in an iota and used as an enum to mark
// unique states
type ID int64

// IDs is a slice of ID
type IDs []ID

// State is an interface that describes one part of a fsm. It is identified by
// a current ID, and can be queried for a slice of IDs that are available to
// transition to. The Transition method of this type should not guard against
// incorrect transitions; rather let that logic be done by the fsm.Transition
// function, which will emit an error if the transition cannot be carried out.
type State interface {
	Current() ID
	Available() IDs
	Transition(ID) State
}

// Transition attempts to transition a state to a new state supplied by the ID,
// if that transition itself is acceptable by the current state.
func Transition(s State, target ID) (State, error) {
	if CanTransition(s, target) {
		return s.Transition(target), nil
	}
	return s, TransitionError{s.Current(), target}
}

// CanTransition determines whether a given transition is acceptable by looking
// through the slice of available IDs on the current state.
func CanTransition(s State, target ID) bool {
	for _, available := range s.Available() {
		if target == available {
			return true
		}
	}

	return false
}

// TransitionError is the error type given off if a transition cannot be carried
// out.
type TransitionError struct {
	From, To ID
}

func (t TransitionError) Error() string {
	return fmt.Sprintf("could not transition from state %d to %d", t.From, t.To)
}
