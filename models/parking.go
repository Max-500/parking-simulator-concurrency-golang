package models

type Parking struct {
	nSpaces int
    spaces *[20]bool
}

func NewParking() *Parking {
	return &Parking{
		nSpaces: 20,
        spaces: &[20]bool{
            true, true, true, true, true, true, true, true, true, true,
            true, true, true, true, true, true, true, true, true, true,
        },
	}
}

func (p *Parking) FindSpaces() int {
    for i, space := range p.spaces {
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

func (p *Parking) GetSpaces() int {
    return p.nSpaces
}

func (p *Parking) GetAllSpaces() *[20]bool {
    return p.spaces
}