package main

import "fmt"

func Nikolay() (a, b int) {
	defer func() {
		a++
		b++
	}()

	return 10, 20
}

func main() {

	fmt.Println(Nikolay())
}
