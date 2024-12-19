package main

import (
	"fmt"
)

func main() {

	var i1, i3 interface{}
	i1 = 42
	i3 = 42
	fmt.Println(i1 == i3)

	var i5 interface{} = []int{1, 2, 3}
	var i7 interface{} = []int{1, 2, 3}

	fmt.Println(i5 == i7)

}
