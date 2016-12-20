package main

import (
	"fmt"
)

func main() {

	channel := make(chan int)
	//channel <- 10 - DeadLock

	go func() {
		channel <- 10
	}()

	fmt.Println(<-channel)
}
