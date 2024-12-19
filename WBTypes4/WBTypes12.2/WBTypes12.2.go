package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "message from ch4"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- "message from ch3"
	}()

	select {
	case msg1 := <-ch1:

		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:

		fmt.Println("Received:", msg2)
	case msg3 := <-ch3:

		fmt.Println("Received:", msg3)
	case msg4 := <-ch4:
		fmt.Println("Received:", msg4)
	}
}
