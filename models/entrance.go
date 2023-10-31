package models

type Entrance struct {
	states [3]string
	actualState string
}

func NewEntrance () *Entrance {
	return &Entrance{
		states: [3]string {"Entering", "Exiting", "Idle"},
		actualState: "Idle",
	}
}

func (e *Entrance) GetState() string {
	return e.actualState
}

func (e *Entrance) SetState(n int) {
	e.actualState = e.states[n]
}