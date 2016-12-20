package main

import (
	"fmt"
	"encoding/json"
)

type Car struct {
	Name string
	Year int
	Color string
}

func main() {
	var car Car

	data := []byte(`{"Name":"Siena", "Year":2017, "Color":"Preto"}`)

	json.Unmarshal(data, &car)

	fmt.Println(car.Name, car.Year, car.Color)
}