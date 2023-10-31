package controllers

import (
	"parking-simulator/models"
	"parking-simulator/views"
	"sync"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type CarController struct {
	model *models.Car
	view *views.CarView
	mu *sync.Mutex
}

func NewCarController(win *pixelgl.Window, mu *sync.Mutex) *CarController {
	return &CarController{
		model: models.NewCar(),
		view: views.NewCarView(win),
		mu: mu,
	}
}

func (cc *CarController) GenerateCars(n int, chCar *chan models.Car) {
	cc.model.GenerateCars(n, *chCar)
}

func (cc *CarController) LoadSprite() {
	cc.view.SetSprite()
}

func (cc *CarController) PaintCar(pos pixel.Vec) {
	cc.view.PaintCar(pos)
}