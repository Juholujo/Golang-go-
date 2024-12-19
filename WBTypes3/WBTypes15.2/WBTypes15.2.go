package main

import (
	"fmt"
)

func main() {
	var numbers []*int
	for _, value := range []int{10, 20, 30, 40} {
		numbers = append(numbers, &value)
	}

	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}
}
