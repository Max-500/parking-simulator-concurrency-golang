package controllers

import (
	"fmt"
	_ "fmt"
	"parking-simulator/models"
	"parking-simulator/views"

	"github.com/faiface/pixel/pixelgl"
)

type ParkingController struct {
    model *models.Parking
    view *views.ParkingView
}

func NewParkingController(win *pixelgl.Window) *ParkingController {
    return &ParkingController{
        model: models.NewParking(),
        view: views.NewParkingView(win),
    }
}

func (pc *ParkingController) PaintParking() {
    pc.view.PaintParking()
}

func (pc *ParkingController) PaintStreet() {
    pc.view.PaintStreet()
}

func (pc *ParkingController) Park(chCar *chan models.Car, entranceController *EntranceController, carController *CarController) {
    for {
        select {
        case _, ok := <-*chCar:
            if !ok{
                return
            }
            pos := pc.model.FindSpaces()
            if pos != -1 {
                if entranceController.model.GetState() == "Idle" || entranceController.model.GetState() == "Entering" {
                    pc.model.ChangeSpace(-1)
                    coordinates := pc.view.GetCoordinates(pos)
                    carController.PaintCar(*coordinates)
                }else{

                }
            }else {
                fmt.Println("Ya no hay espacios disponibles")
            }
        }
    }
}
