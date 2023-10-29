package views

import (
	"parking-simulator/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ParkingView struct {
	win *pixelgl.Window
	utils *utils.Utils
}

func NewParkingView(win *pixelgl.Window) *ParkingView {
	return &ParkingView{
		win: win,
	}
}

func (pw *ParkingView) PaintParking() {
	picParking, err := pw.utils.LoadPicture("./assets/parking.png")
	if err != nil {
		panic(err)
	}

	parking := pw.utils.NewSprite(picParking, picParking.Bounds())

	matrix := pixel.IM
	matrix = pixel.IM.Moved(pixel.V(512, 469))
	parking.Draw(pw.win, matrix)
}

func (pw *ParkingView) PaintStreet() {
	picStreet, err := pw.utils.LoadPicture("./assets/street.png")
	if err != nil {
		panic(err)
	}

	street := pw.utils.NewSprite(picStreet, picStreet.Bounds())

	street.Draw(pw.win, pixel.IM.Moved(pixel.V(512, 85)))
}