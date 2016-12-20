package main


import "fmt"

func main() {

	m := make(map[string]int)

	m["a"] = 7
	m["b"] = 77
	m["c"] = 777

	fmt.Println(m)

	delete(m, "c")

	fmt.Println(m)
	fmt.Println(m["c"])

	_, exists := m["c"]

	fmt.Println(exists)

	value, exists := m["a"]

	fmt.Println(value)


	//var x = map[string]int{}
	//
	//fmt.Println(x)


	// sem o make
	newMap := map[string]int{"x":5, "y":10}
	fmt.Println(newMap)

	if value, exists := m["p"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("Não existe")
	}
}