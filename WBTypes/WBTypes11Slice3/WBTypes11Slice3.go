package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := make([]int, 2, 4) // длина 2, емкость 4
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("Aftre make: %v len: %d cap: %d ptr: %x\n", slice, len(slice), cap(slice), header.Data)
	slice[0] = 1
	slice[1] = 2

	// Получаем адрес базового массива
	header = (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("Before append: %v len: %d cap: %d ptr: %x\n", slice, len(slice), cap(slice), header.Data)

	slice = append(slice, 3, 4) // добавление элементов в слайс

	// Получаем новый адрес базового массива
	header = (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("After append: %v len: %d cap: %d ptr: %x\n", slice, len(slice), cap(slice), header.Data)

	slice = append(slice, 5) // добавление еще одного элемента

	// Получаем новый адрес базового массива
	header = (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("After another append: %v len: %d cap: %d ptr: %x\n", slice, len(slice), cap(slice), header.Data)
}
