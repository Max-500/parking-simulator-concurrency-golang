package views

import (
	"parking-simulator/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type CarView struct {
	win   *pixelgl.Window
	utils *utils.Utils
	sprite *pixel.Sprite
}

type SpriteCar struct {
	img *pixel.Sprite
	Id int
}

func NewCarView(win *pixelgl.Window) *CarView {
	return &CarView{
		win: win,
	}
}

func (cw *CarView) SetSprite() {
	picCar, _ := cw.utils.LoadPicture("./assets/car.png")
	cw.sprite = cw.utils.NewSprite(picCar, picCar.Bounds())
}

func (cw *CarView) PaintCar(pos pixel.Vec) *pixel.Sprite {
	cw.sprite.Draw(cw.win, pixel.IM.Moved(pos))
	return cw.sprite
}

func NewImgCar(spr *pixel.Sprite, Id int) *SpriteCar {
	return &SpriteCar{
		img: spr,
		Id: Id,
	}
}