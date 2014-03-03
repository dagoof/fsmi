package fsm

import (
	"fmt"
)

type State int64
type States []State

type Machine interface {
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
