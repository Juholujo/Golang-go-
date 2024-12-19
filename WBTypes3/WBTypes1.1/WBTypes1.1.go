package main

import "fmt"

func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Привет")
	testSlice = append(testSlice, "Привет")
	test(testSlice)

	fmt.Println(testSlice) //

}

func test(testSlice []string) {
	testSlice = append(testSlice, "WB")

}

//fmt.Println("ANDREW", testSlice)
