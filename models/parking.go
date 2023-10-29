package models

import "fmt"

type Parking struct {
    msg string
}

func NewParking(msg string) *Parking {
    return &Parking{
        msg: msg,
    }
}

func (p *Parking) MiFuncion() {
    fmt.Println("MiFuncion en el modelo Parking")
}