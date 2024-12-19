package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type emptyStruct struct{}

func main() {

	var Struct1 = struct{}{}
	var Struct2 = struct{}{}
	var emptyInstance = emptyStruct{}
	var byteArray = [0]byte{}

	// Вывод указателей на переменные
	fmt.Println("Pointer to Struct1:", unsafe.Pointer(&Struct1))
	fmt.Println("Pointer to Struct2:", unsafe.Pointer(&Struct2))
	fmt.Println("Pointer to emptyInstance:", unsafe.Pointer(&emptyInstance))
	fmt.Println("Pointer to byteArray:", unsafe.Pointer(&byteArray))

	// Вывод размеров переменных
	fmt.Println("Size of Struct1:", unsafe.Sizeof(Struct1))
	fmt.Println("Size of Struct2:", unsafe.Sizeof(Struct2))
	fmt.Println("Size of emptyInstance:", unsafe.Sizeof(emptyInstance))
	fmt.Println("Size of byteArray:", unsafe.Sizeof(byteArray))

	fmt.Println("Size of Struct1 (reflect):", reflect.TypeOf(Struct1).Size())
	fmt.Println("Size of Struct2 (reflect):", reflect.TypeOf(Struct2).Size())
	fmt.Println("Size of emptyInstance (reflect):", reflect.TypeOf(emptyInstance).Size())
	fmt.Println("Size of byteArray (reflect):", reflect.TypeOf(byteArray).Size())
}
