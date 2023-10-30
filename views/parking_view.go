package views

import (
	"parking-simulator/utils"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type ParkingView struct {
	win *pixelgl.Window
	utils *utils.Utils
	spaces [20]pixel.Vec
}

func NewParkingView(win *pixelgl.Window) *ParkingView {
	return &ParkingView{
		win: win,
		spaces: [20]pixel.Vec{
			pixel.V(835, 300),
			pixel.V(835, 385),
			pixel.V(835, 465),
			pixel.V(835, 545),
			pixel.V(835, 625),

			pixel.V(685, 625),
			pixel.V(685, 545),
			pixel.V(685, 465),
			pixel.V(685, 385),
			pixel.V(685, 300),

			pixel.V(175, 625),
			pixel.V(175, 545),
			pixel.V(175, 465),
			pixel.V(175, 385),
			pixel.V(175, 300),

			pixel.V(325, 300),
			pixel.V(325, 385),
			pixel.V(325, 465),
			pixel.V(325, 545),
			pixel.V(325, 625),
		},
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

func (pw *ParkingView) GetCoordinates(n int) *pixel.Vec {
	return &pw.spaces[n+1]
}