package main

import "fmt"

func got() int {
	return 20
}

func gotPanic() int {
	panic("No panic, Yura! ")
	return 30
}

func main() {
	first := make(chan int, 1)
	first <- 10

	second := make(chan int, 1)

	select {
	case second <- got():
		fmt.Println("SEND ME!")
	case first <- gotPanic():
	}
}
