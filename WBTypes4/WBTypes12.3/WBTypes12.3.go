package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// Значения для case вычисляются заранее
	val1 := expensiveCalculation("case 1")
	val2 := expensiveCalculation("case 2")

	select {
	case ch <- val1: // Попытка записать val1 в канал
		fmt.Println("Sent value:", val1)
	case ch <- val2: // Попытка записать val2 в канал
		fmt.Println("Sent value:", val2)
	default:
		fmt.Println("No channels ready, but calculations are already done")
	}
}

func expensiveCalculation(name string) int {
	fmt.Println("Calculating for", name)
	return len(name)
}
