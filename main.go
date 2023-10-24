package main

import (
	_ "fmt"
	"image"
	_"math/rand"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	// ancho, largo
	cfg := pixelgl.WindowConfig{
		Title:  "Parking Simulator!!!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	picParking, err := loadPicture("./assets/parking3.png")
	if err != nil {
		panic(err)
	}

	picStreet, err := loadPicture("./assets/street3.png")
	if err != nil {
		panic(err)
	}


	picCar, err := loadPicture("./assets/car.png")
	if err != nil {
		panic(err)
	}

	picEntranceClose, err := loadPicture("./assets/open-entrance1.png")
	if err != nil {
		panic(err)
	}

	parking := pixel.NewSprite(picParking, picParking.Bounds())
	street := pixel.NewSprite(picStreet, picStreet.Bounds())
	car := pixel.NewSprite(picCar, picCar.Bounds())
	closeEntrance := pixel.NewSprite(picEntranceClose, picEntranceClose.Bounds())

	for !win.Closed() {
		win.Clear(colornames.Black)

		matrix := pixel.IM
		matrix = pixel.IM.Moved(pixel.V(512, 469))
		parking.Draw(win, matrix)


		street.Draw(win, pixel.IM.Moved(pixel.V(512, 85)))
		car.Draw(win, pixel.IM.Moved(pixel.V(112, 469)))
		closeEntrance.Draw(win, pixel.IM.Moved(pixel.V(920, 200)))

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
