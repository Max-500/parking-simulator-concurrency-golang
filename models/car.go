package models

import (
	"fmt"
	"math/rand"
	"parking-simulator/utils"
	"sync"
	"time"

	"github.com/faiface/pixel"
)

type Car struct {
	ParkingTime int // Tiempo en segundos que estará estacionado
	Id int
}// 10 + 15 (20)

func NewCar() *Car {
	rand.Seed(time.Now().UnixNano()) // Inicializar el generador de números aleatorios con una semilla única
	parkingTime := rand.Intn(10) + 15 // Generar un número aleatorio
	return &Car{ParkingTime: parkingTime}
}

func (c *Car) GenerateCars(n int, ch chan Car) {
	for i := 1; i <= n; i++ {
		car := NewCar()
		car.Id = i
		ch<- *car
		rand.Seed(time.Now().UnixNano()) 
		newTime := rand.Intn(2) + 1
		time.Sleep(time.Second * time.Duration(newTime))
	}
	close(ch)
	fmt.Println("Se termino de generar los autos")
}

func (c *Car) Timer(pos int, pc *Parking, mu *sync.Mutex, spaces *[20]bool, chEntrance *chan int, sprite *pixel.Sprite, chWin chan utils.ImgCar, coo pixel.Vec) {
	mu.Lock()
	data := utils.NewImgCar(sprite, pos, true, coo)
	chWin<-*data
	*chEntrance<-0
	mu.Unlock()

	mu.Lock()
	pc.nSpaces--
	fmt.Println("El auto", c, "acaba de estacionarse y esta estacionado en el lugar número:", pos)
	fmt.Println("Quedan", pc.nSpaces, " espacios disponibles")
	mu.Unlock()

	time.Sleep(time.Second * time.Duration(c.ParkingTime))

	fmt.Println("El auto", c, c.Id,"acaba de saliry estaba estacionado en el lugar número:", pos)
	
	mu.Lock()
	data = utils.NewImgCar(sprite, pos, false, coo)
	chWin<-*data
	pc.nSpaces = pc.nSpaces + 1
	spaces[pos] = true
	fmt.Println("Quedan", pc.nSpaces, " espacios disponibles, despues de que se fuera el auto")
	mu.Unlock()

	mu.Lock()
	*chEntrance<-1
	mu.Unlock()
}