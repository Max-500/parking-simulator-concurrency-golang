package views

import (
	"parking-simulator/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type EntranceView struct {
	win *pixelgl.Window
	utils *utils.Utils
	states [3]pixel.Sprite
}

func NewEntranceView(win *pixelgl.Window) *EntranceView {
	return &EntranceView{
		win: win,
	}
}

func (ew *EntranceView) LoadStatesImages() [3]pixel.Sprite {
	picEntranceOpen, _ := ew.utils.LoadPicture("./assets/open-entrance.png")
	picEntranceOpening, _ := ew.utils.LoadPicture("./assets/opening-entrance.png")
	picEntranceClose, _ := ew.utils.LoadPicture("./assets/close-entrance.png")
	
	openEntrance := ew.utils.NewSprite(picEntranceOpen, picEntranceOpen.Bounds())
	openingEntrance := ew.utils.NewSprite(picEntranceOpening, picEntranceOpening.Bounds())
	closeEntrance := ew.utils.NewSprite(picEntranceClose, picEntranceClose.Bounds())

	return [3]pixel.Sprite{*openEntrance, *openingEntrance, *closeEntrance}
}

func (ew *EntranceView) SetStateImages(imgs [3]pixel.Sprite) {
	ew.states = imgs
}

func (ew *EntranceView) PainEntrance(img int) {
	ew.states[img].Draw(ew.win, pixel.IM.Moved(pixel.V(920, 200)))
}