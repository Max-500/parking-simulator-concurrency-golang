package controllers

import (
	"parking-simulator/models"
	"parking-simulator/views"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type CarController struct {
	model *models.Car
	view *views.CarView
}

func NewCarController(win *pixelgl.Window) *CarController {
	return &CarController{
		model: models.NewCar(),
		view: views.NewCarView(win),
	}
}

func (cc *CarController) GenerateCars(n int, chCar *chan models.Car) {
	cc.model.GenerateCars(n, *chCar)
}

func (cc *CarController) LoadSprite () {
	cc.view.SetSprite()
}

func (cc *CarController) PaintCar(pos pixel.Vec) {
	cc.view.PaintCar(pos)
}