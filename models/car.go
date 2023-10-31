package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Car struct {
	ParkingTime int // Tiempo en segundos que estará estacionado
	Id int
}

func NewCar() *Car {
	rand.Seed(time.Now().UnixNano()) // Inicializar el generador de números aleatorios con una semilla única
	parkingTime := rand.Intn(10) + 1 // Generar un número aleatorio entre 1 y 5 segundos
	return &Car{ParkingTime: parkingTime}
}

func (c *Car) GenerateCars(n int, ch chan Car) {
	for i := 1; i <= n; i++ {
		car := NewCar()
		car.Id = i
		ch<- *car
		rand.Seed(time.Now().UnixNano()) 
		newTime := rand.Intn(5) + 1
		time.Sleep(time.Second * time.Duration(newTime))
	}
	close(ch)
	fmt.Println("Se termino de generar los autos")
}

func (c *Car) Timer(pos int, pc *Parking, mu *sync.Mutex, spaces *[20]bool, chEntrance *chan int) {
	mu.Lock()
	*chEntrance<-0
	mu.Unlock()

	mu.Lock()
	pc.nSpaces--
	fmt.Println("El auto", c, "con el ID", c.Id,"acaba de estacionarse")
	fmt.Println("Quedan", pc.nSpaces, " espacios disponibles")
	mu.Unlock()

	time.Sleep(time.Second * time.Duration(c.ParkingTime))

	fmt.Println("El auto", c, "con el ID", c.Id,"acaba de salir")
	mu.Lock()
	pc.nSpaces = pc.nSpaces + 1
	spaces[pos] = true
	fmt.Println("Quedan", pc.nSpaces, " espacios disponibles")
	mu.Unlock()

	mu.Lock()
	*chEntrance<-1
	mu.Unlock()
}