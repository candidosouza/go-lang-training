package main

import "fmt"

type vehicle interface {
	start() string
}

type car struct {
	name string
}

func(c car) start() string {
	return "O carro " + c.name + " ligou"
}

type motorcycle struct {
	name string
}

func(mc motorcycle) start() string {
	return "O Moto " + mc.name + " ligou"
}

func startVehicle(v vehicle) string {
	return v.start()
}

func main() {
	c := car{"Siena"}
	mc := motorcycle{"BMW"}

	//fmt.Println(c.startCar())
	//fmt.Println(mc.startMc())

	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(mc))
}
