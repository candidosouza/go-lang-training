package main

import "fmt"

func main() {

	msg := make(chan string)

	go func() {
		msg <- "Hello Word"
	}()

	result := <- msg

	fmt.Println(result)

}
