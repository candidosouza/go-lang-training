package main

import (
	"fmt"
)

type Car struct {
	Name string
	Year int
	Color string
}

type ElectricCar struct {
	Car
	Name string
	Electric bool
}

func (c Car) info() string  {
	return fmt.Sprintf(" Carro: %s\n Ano: %i\n Cor: %d\n", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Siena", 2017, "Preto"}

	//electricCar := ElectricCar{ Car: Car{"Fusca", 69, "Branco"}, Electric: true, Name: "Ferrari"}
	electricCar := ElectricCar{ Car: car1, Electric: true, Name: "Ferrari"}

	fmt.Println(car1.info())
	fmt.Println(electricCar)
	fmt.Println(electricCar.Name)
	fmt.Println(electricCar.Car.Name)
}