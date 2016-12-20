package main

import (
	"fmt"
)

func main() {
	numbers := generate(2, 4, 6)

	result := divide(numbers)

	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, numbers := range numbers {
			channel <- numbers
		}
	}()

	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for numbers := range input {
			channel <- numbers / 2
		}

		close(channel)
	}()

	return channel
}
