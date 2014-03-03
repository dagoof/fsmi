package fsm

import (
	"fmt"
)

type State int64
type States []State

type MachineMaker func() (Machine, error)

type Machine interface {
	Set(State)
	Current() State
	ToStates() States
	TransitionTo(State) MachineMaker
}

func Transition(m Machine, target State) (Machine, error) {
	if CanTransition(m, target) {
		return m.TransitionTo(target)()
	}
	return m, TransitionError{m.Current(), target}
}


func Self(m Machine, target State) MachineMaker {
	return func() (Machine, error) {
		if CanTransition(m, target) {
			m.Set(target)
			return m, nil
		}
		return m, TransitionError{m.Current(), target}
	}
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
	return fmt.Sprintf("could not transition from %s to %s", t.From, t.To)
}
