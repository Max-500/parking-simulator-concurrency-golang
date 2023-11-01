package main

import (
	_ "image/png"
	"parking-simulator/controllers"
	"parking-simulator/models"
	"parking-simulator/utils"
	"sync"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Parking Simulator!!!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Canales
	carChannel := make(chan models.Car, 100)
	entranceChannel := make(chan int)
	winChannel := make(chan utils.ImgCar)

	// Semaforo para coordinar
	mu := &sync.Mutex{}

	// Controladores
	parkingController := controllers.NewParkingController(win, mu)
	entranceController := controllers.NewEntranceController(win, mu)
	carController := controllers.NewCarController(win, mu)

	carController.LoadSprite()


	go parkingController.Park(&carChannel, entranceController, carController, &entranceChannel, winChannel)
	entranceController.LoadStates()
	go carController.GenerateCars(100, &carChannel)

	var arr []utils.ImgCar
	for !win.Closed() {
		win.Clear(colornames.Black)

		parkingController.PaintParking()
		parkingController.PaintStreet()

		select {
		case val := <-winChannel:

			if val.GetStatus() {
				arr = append(arr, val)
			}else{
				var arrAux []utils.ImgCar
				for _, value := range arr {
					if value.GetId() != val.GetId() {
						arrAux = append(arrAux, value)
					}
				}
				arr = arr[:0]
				arr = append(arr, arrAux...)
			}
		}

		for _, value := range arr {
			sprite := value.GetSprite()
			pos := value.GetPos()
			sprite.Draw(win, pixel.IM.Moved(pos))
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}