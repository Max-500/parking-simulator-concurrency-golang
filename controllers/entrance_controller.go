package controllers

import (
	"parking-simulator/models"
	"parking-simulator/views"
	"github.com/faiface/pixel/pixelgl"
)

type EntranceController struct {
	model *models.Entrance
	view *views.EntranceView
}

func NewEntranceController (win *pixelgl.Window) *EntranceController {
	return &EntranceController{
		model: models.NewEntrance(),
		view: views.NewEntranceView(win),
	}
}

func (ec *EntranceController) LoadStates() {
	imgs := ec.view.LoadStatesImages()
	ec.view.SetStateImages(imgs)
}

func (ec * EntranceController) PaintEntrance(pos int) {
	ec.view.PainEntrance(pos)
}