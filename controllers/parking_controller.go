package controllers

import (
	"parking-simulator/models"
	"parking-simulator/views"
	"github.com/faiface/pixel/pixelgl"
)

type ParkingController struct {
    model *models.Parking
    view *views.ParkingView
}

func NewParkingController(win *pixelgl.Window) *ParkingController {
    return &ParkingController{
        model: models.NewParking("Hello, World!"),
        view: views.NewParkingView(win),
    }
}

func (pc *ParkingController) LlamarMiFuncion() {
    pc.model.MiFuncion()
}

func (pc *ParkingController) PaintParking() {
    pc.view.PaintParking()
}

func (pc *ParkingController) PaintStreet() {
    pc.view.PaintStreet()
}