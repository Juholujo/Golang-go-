package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "message from ch2"
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		default:
			fmt.Println("No message received, performing default action")
			time.Sleep(1 * time.Second)
		}
	}
}
