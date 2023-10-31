package controllers

import (
	"parking-simulator/models"
	"parking-simulator/utils"
	"parking-simulator/views"
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

type ParkingController struct {
    model *models.Parking
    view *views.ParkingView
    mu *sync.Mutex
}

func NewParkingController(win *pixelgl.Window, mu *sync.Mutex) *ParkingController {
    return &ParkingController{
        model: models.NewParking(),
        view: views.NewParkingView(win),
        mu: mu,
    }
}

func (pc *ParkingController) PaintParking() {
    pc.view.PaintParking()
}

func (pc *ParkingController) PaintStreet() {
    pc.view.PaintStreet()
}

func (pc *ParkingController) Park(chCar *chan models.Car, entranceController *EntranceController, carController *CarController, chEntrance *chan int, chWin chan utils.ImgCar) {
    go pc.ChangingState(chEntrance, entranceController)
    for {
        select {
        case car, ok := <-*chCar:
            if !ok{
                return
            }
            pos := pc.model.FindSpaces()
            if pos != -1 {
                coo := pc.view.GetCoordinates(pos)
                carController.view.SetSprite()
                sprite := carController.view.PaintCar(coo)
                if entranceController.model.GetState() == "Idle" || entranceController.model.GetState() == "Entering" {
                    go car.Timer(pos, pc.model, pc.mu, pc.model.GetAllSpaces(), chEntrance, sprite, chWin, coo)
                }else{
                    *chEntrance<-0
                    go car.Timer(pos, pc.model, pc.mu, pc.model.GetAllSpaces(), chEntrance, sprite, chWin, coo)
                }
            }
        }
    }
}

func (pc *ParkingController) ChangingState(chEntrance *chan int, entrancecontroller *EntranceController) {
    for {
        select{
        case change, _ := <-*chEntrance:
            entrancecontroller.model.SetState(change)
        }
    }
}