package main

import (
	"image"
	_ "image/png"
	"os"
	"parking-simulator/controllers"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// "./assets/car.png"

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
	cfg := pixelgl.WindowConfig{
		Title:  "Parking Simulator!!!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	picEntranceClose, err := loadPicture("./assets/open-entrance.png")
	if err != nil {
		panic(err)
	}

	closeEntrance := pixel.NewSprite(picEntranceClose, picEntranceClose.Bounds())

	parkingController := controllers.NewParkingController(win)
	parkingController.LlamarMiFuncion()

	entranceController := controllers.NewEntranceController(win)
	entranceController.LoadStates()

	
	for !win.Closed() {
		win.Clear(colornames.Black)

		parkingController.PaintParking()
		parkingController.PaintStreet()

		closeEntrance.Draw(win, pixel.IM.Moved(pixel.V(920, 200)))

		win.Update()
	}

}

func main() {
	pixelgl.Run(run)
}