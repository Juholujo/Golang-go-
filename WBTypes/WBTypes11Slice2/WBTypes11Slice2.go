package main

import "fmt"

func main() {

	var slice []int // nil
	fmt.Println(slice)
	fmt.Println("Действительно ли slice = nil:", slice == nil)

	slice = []int{}
	fmt.Println(slice)
	fmt.Println("А сейчас slice = nil?", slice == nil)

	/*Можно проверять на пустоту slice через len*/
	//var slice []int
	//
	//fmt.Println("Действительно ли slice = nil:", len(slice) == 0)
	//
	//slice = []int{}
	//
	//fmt.Println("А сейчас slice = nil?", len(slice) == 0)

}
