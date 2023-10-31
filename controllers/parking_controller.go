package controllers

import (
	"fmt"
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
    ch *chan utils.CarSprite
}

func NewParkingController(win *pixelgl.Window, mu *sync.Mutex, ch chan utils.CarSprite) *ParkingController {
    return &ParkingController{
        model: models.NewParking(),
        view: views.NewParkingView(win),
        mu: mu,
        ch: &ch,
    }
}

func (pc *ParkingController) PaintParking() {
    pc.view.PaintParking()
}

func (pc *ParkingController) PaintStreet() {
    pc.view.PaintStreet()
}

func (pc *ParkingController) Park(chCar *chan models.Car, entranceController *EntranceController, carController *CarController, chEntrance *chan int) {
    go pc.ChangingState(chEntrance, entranceController)
    for {
        select {
        case car, ok := <-*chCar:
            if !ok{
                return
            }
            pos := pc.model.FindSpaces()
            if pos != -1 {
                if entranceController.model.GetState() == "Idle" || entranceController.model.GetState() == "Entering" {
                    //sprite := carController.view.PaintCar(coo)
                    //coo := pc.view.GetCoordinates(pos)
                    go car.Timer(pos, pc.model, pc.mu, pc.model.GetAllSpaces(), chEntrance)
                }else{
                    fmt.Println("Me atoro")
                    *chEntrance<-0
                    go car.Timer(pos, pc.model, pc.mu, pc.model.GetAllSpaces(), chEntrance)
                }
            }else {
                fmt.Println("Se lleno el estacionamiento")
                select{
                    // Aqui esperar que algun canal me confirme que ya hay espacio
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