package main

import "fmt"

func somePanic() {
	panic("Паника в somePanic()")
}
func main() {
	fmt.Println("Начало фунцкии main")

	somePanic()

	recError := recover()
	fmt.Println(recError)
	fmt.Println("Завершение фунцкии main")
}
