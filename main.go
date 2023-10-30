package main

import (
	_"fmt"
	_ "image/png"
	"parking-simulator/controllers"
	"parking-simulator/models"
	"time"
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
	carChannel := make(chan models.Car)

	// Controladores
	parkingController := controllers.NewParkingController(win)
	entranceController := controllers.NewEntranceController(win)
	carController := controllers.NewCarController(win)

	carController.LoadSprite()


	go parkingController.Park(&carChannel, entranceController, carController)
	entranceController.LoadStates()
	go carController.GenerateCars(30, &carChannel)

	imageChangeChannel := make(chan int)

	go func() {
		a := 0
		for {
			imageChangeChannel <- a
			a = (a + 1) % 2 // Cambia el valor entre 0 y 1
			time.Sleep(2 * time.Second)
		}
	}()

	for !win.Closed() {
		win.Clear(colornames.Black)

		parkingController.PaintParking()
		parkingController.PaintStreet()

		select {
		case value := <-imageChangeChannel:
			entranceController.PaintEntrance(value)
		default:
			// No se ha enviado un nuevo valor, usa el valor actual
			entranceController.PaintEntrance(2)
		}

		win.Update()
		time.Sleep(time.Second * 5)
	}
}

func main() {
	pixelgl.Run(run)
}