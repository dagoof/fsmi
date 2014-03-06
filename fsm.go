// Package fsmi exposes a method of working with FSM's through interfaces rather
// than concrete state structs. By satisfying the State interface and using the
// Transition function, it is possible to have arbitrary types act behave as FSM
// states without embedding State structs everywhere and cluttering your types.
package fsmi

import (
	"fmt"
)

// ID is a type meant to be declared in an iota and used as an enum to mark
// states. Implements the Identifier interface.
type ID uint64

// Identity returns the value of the ID as a value that can be  used to identify
// states.
func (id ID) Identity() uint64 {
	return uint64(id)
}

// Identifier describes a type that can Identify itself with a uint64. This is
// used to signal transitions to a state, and to allow a state to express which
// transitions are available.
type Identifier interface {
	Identity() uint64
}

// State is an interface that describes one part of a fsm. It is identified by
// an Identifier, and can be queried for a slice of Identifiers that are
// available to transition to. The Transition method of this type should not
// guard against incorrect transitions; rather let that logic be done by the
// fsm.Transition function, which will emit an error if the transition
// cannot be carried out.
type State interface {
	Current() Identifier
	Available() []Identifier
	Transition(Identifier) State
}

// Transition attempts to transition a state to a new state supplied by the
// Identifier, if that transition itself is acceptable by the current state.
func Transition(s State, target Identifier) (State, error) {
	if CanTransition(s, target) {
		return s.Transition(target), nil
	}
	return s, TransitionError{s.Current(), target}
}

// CanTransition determines whether a given transition is acceptable by looking
// through the slice of available Identifiers on the current state.
func CanTransition(s State, target Identifier) bool {
	targetIdentity := target.Identity()

	for _, available := range s.Available() {
		if targetIdentity == available.Identity() {
			return true
		}
	}

	return false
}

// TransitionError is the error type given off if a transition cannot be carried
// out.
type TransitionError struct {
	From Identifier
	To   Identifier
}

func (t TransitionError) Error() string {
	return fmt.Sprintf(
		"could not transition from state %d to %d",
		t.From.Identity(), t.To.Identity(),
	)
}
