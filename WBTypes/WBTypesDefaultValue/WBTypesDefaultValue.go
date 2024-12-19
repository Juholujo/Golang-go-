package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string
	var num int
	var boolean bool
	var slice []int
	var m map[string]int
	var iface interface{}

	//fmt.Println("Default string:", str)      // ""
	//fmt.Println("Default int:", num)         // 0
	//fmt.Println("Default bool:", boolean)    // false
	//fmt.Println("Default slice:", slice)     // nil
	//fmt.Println("Default map:", m)           // nil
	//fmt.Println("Default interface:", iface) // nil

	fmt.Printf("Default string: %v\n", str)      // ""
	fmt.Printf("Default int: %v\n", num)         // 0
	fmt.Printf("Default bool: %v\n", boolean)    // false
	fmt.Printf("Default slice: %v\n", slice)     // nil
	fmt.Printf("Default map: %v\n", m)           // nil
	fmt.Printf("Default interface: %v\n", iface) // nil

	// Check slice
	fmt.Println("Slice is nil:", slice == nil)                  // true
	fmt.Println("Slice reflect type:", reflect.TypeOf(slice))   // []int
	fmt.Println("Slice reflect value:", reflect.ValueOf(slice)) // <nil>

	// Check map
	fmt.Println("Map is nil:", m == nil)                  // true
	fmt.Println("Map reflect type:", reflect.TypeOf(m))   // map[string]int
	fmt.Println("Map reflect value:", reflect.ValueOf(m)) // <nil>

	// Check interface
	fmt.Println("Interface is nil:", iface == nil)                  // true
	fmt.Println("Interface reflect type:", reflect.TypeOf(iface))   // <nil>
	fmt.Println("Interface reflect value:", reflect.ValueOf(iface)) // <nil>
}
