package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := make([]int, 4, 7)
	array := [4]int(slice[:4])

	slice[0] = 7

	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	fmt.Println("VALUE slice:", reflect.ValueOf(slice))
	fmt.Println("TYPE slice:", reflect.TypeOf(slice))

	fmt.Printf("array: %v, len: %d, cap: %d\n", array, len(array), cap(array))

	fmt.Println("VALUE array:", reflect.ValueOf(array))
	fmt.Println("TYPE array:", reflect.TypeOf(array))
}

/*В первом примере массив array является указателем на часть среза, и изменения в срезе могут повлиять на массив.
Во втором примере массив array создается как новый массив, и изменения в срезе не влияют на массив.*/
