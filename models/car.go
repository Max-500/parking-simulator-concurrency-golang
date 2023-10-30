package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct {
	ParkingTime int // Tiempo en segundos que estará estacionado
	Id int
}

func NewCar() *Car {
	rand.Seed(time.Now().UnixNano()) // Inicializar el generador de números aleatorios con una semilla única
	parkingTime := rand.Intn(5) + 1 // Generar un número aleatorio entre 1 y 5 segundos
	return &Car{ParkingTime: parkingTime}
}

func (c *Car) GenerateCars(n int, ch chan Car) {
	for i := 1; i <= n; i++ {
		car := NewCar()
		car.Id = i
		ch<- *car
		fmt.Println(car.Id)
		rand.Seed(time.Now().UnixNano()) 
		newTime := rand.Intn(5) + 1
		time.Sleep(time.Second * time.Duration(newTime))

	}
	close(ch)
	fmt.Println("Se termino de generar los autos")
}
