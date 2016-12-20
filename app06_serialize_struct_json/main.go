package main

import (
	"fmt"
	"encoding/json"
)

type Car struct {
	Name string `json:"Modelo"`
	Year int`json:"-"`
	Color string
}

func main() {
	car1 := Car{"Siena", 2017, "Preto"}

	result, _ := json.Marshal(car1)

	//fmt.Println(result)

	fmt.Println(string(result))
}