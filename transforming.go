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

	for state := range t.To {
		states = append(states, state)
	}
	return states
}

func (t *Transforming) Transition(target State) Machine {
	return t.To[target](t)
}
