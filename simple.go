package fsm

type Simple struct {
	State State
	To    []State
}

func (s *Simple) Current() State {
	return s.State
}

func (s *Simple) ToStates() States {
	return s.To
}

func (s *Simple) Set(target State) {
	s.State = target
}

func (s *Simple) TransitionTo(target State) MachineMaker {
	return Self(s, target)
}

type Transitioning struct {
	State State
	To    map[State]MachineMaker
}
