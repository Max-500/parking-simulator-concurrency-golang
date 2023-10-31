package controllers

import (
	"parking-simulator/models"
	"parking-simulator/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type EntranceController struct {
	model *models.Entrance
	view *views.EntranceView
	mu *sync.Mutex
}

func NewEntranceController (win *pixelgl.Window, mu *sync.Mutex) *EntranceController {
	return &EntranceController{
		model: models.NewEntrance(),
		view: views.NewEntranceView(win),
		mu: mu,
	}
}

func (ec *EntranceController) LoadStates() {
	imgs := ec.view.LoadStatesImages()
	ec.view.SetStateImages(imgs)
}

func (ec * EntranceController) PaintEntrance(pos int) {
	ec.view.PainEntrance(pos)
}