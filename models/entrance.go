package models

type Entrance struct {
	states [3]string
}

func NewEntrance () *Entrance {
	return &Entrance{
		states: [3]string {"Entering", "Exiting", "Idle"},
	}
}