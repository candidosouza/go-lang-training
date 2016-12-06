package main

import "fmt"

func main() {

	// decimal
	fmt.Println(77)

	// binary
	// %d = decimal
	// %b = binary
	fmt.Printf("%d - %b \n", 77, 77)

	// hexadecimal
	// %x = hexadecimal
	fmt.Printf("%d - %b - %x \n", 77, 77, 77)
	fmt.Printf("%d - %b - %#x \n", 77, 77, 77)
	fmt.Printf("%d \t %b \t %#x \n", 77, 77, 77)

	// loop
	for i := 0; i < 200; i ++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
