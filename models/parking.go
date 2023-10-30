package models

import "fmt"

type Parking struct {
	nSpaces int
    spaces [20]bool
}

func NewParking() *Parking {
	return &Parking{
		nSpaces: 20,
        spaces: [20]bool{
            true, true, true, true, true, true, true, true, true, true,
            true, true, true, true, true, true, true, true, true, true,
        },
	}
}

func (p *Parking) FindSpaces() int {
    for i, space := range p.spaces {
        fmt.Println(space, "entro", i)
        if space {
            p.spaces[i] = false
            return i
        }
    }
    return -1
}

func (p *Parking) ChangeSpace(n int) {
	p.nSpaces = p.nSpaces + n
}