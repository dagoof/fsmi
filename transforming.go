package fsm

type Transforming struct {
	State State
	To    map[State]MachineMaker
}

func (t *Transforming) Current() State {
	return t.State
}

func (t *Transforming) ToStates() States {
	var states States

	for _, state := range t.To {
		states = append(states, state)
	}
	return states
}

func (t *Transforming) TransitionTo(target State) MachineMaker {
	return t.To[target]
}
