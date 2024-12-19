package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	// Второй пример:

	list := make([]int, 4, 4)
	fmt.Println("before:", list)

	headerList := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	fmt.Printf("Memory list: Data=0x%x Len=%d Cap=%d\n", headerList.Data, headerList.Len, headerList.Cap)

	list2 := append(list, 1)

	headerList2 := (*reflect.SliceHeader)(unsafe.Pointer(&list2))
	fmt.Printf("Memory list2: Data=0x%x Len=%d Cap=%d\n", headerList2.Data, headerList2.Len, headerList2.Cap)

	list[0] = 4

	list2[0] = 9

	fmt.Println(list)
	fmt.Println(list2)
}
