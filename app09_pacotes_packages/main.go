package main

import (
	"fmt"
	"github.com/GoLangTraning/app09_pacotes_packages/car"
)


func main() {
	car := car.Car{"Siena", "Preto"}
	fmt.Println(car.Start())
}
