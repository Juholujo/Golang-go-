package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//var a map[string]int // nil
	//
	//fmt.Println(a == nil) // true
	//
	//fmt.Println(a["key"]) // 0 (чтение допустимо)
	//
	//a["key"] = 42 // panic: assignment to entry in nil map

	a := make(map[string]int)

	fmt.Println("VALUE p:", reflect.ValueOf(a))

	fmt.Println("TYPE p:", reflect.TypeOf(a))
	fmt.Println("TYPE p:", unsafe.Sizeof(a))

	fmt.Println(a == nil)
	fmt.Println(a["key"]) // 0 (чтение допустимо)
	a["key"] = 42
	fmt.Println(a["key"]) // 42
}
