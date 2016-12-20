package main

import (
	"fmt"
)

type Car struct {
	Name string
	Year int
	Color string
}

func (c Car) info() string  {
	return fmt.Sprintf(" Carro: %s\n Ano: %i\n Cor: %d\n", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Siena", 2017, "Preto"}
	fmt.Println(car1.info())
}