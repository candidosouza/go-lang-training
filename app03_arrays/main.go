package main

import "fmt"

func main() {
	var x [10]int
	fmt.Println(x)
	fmt.Println(len(x))

	x[0] = 7
	x[1] = 10
	x[2] = 13

	fmt.Println(x)

	fmt.Println(x[2])

	var x [10]string
	fmt.Println(x)

	z := [5]int{7,77,10,9,5}
	fmt.Println(z)
}
