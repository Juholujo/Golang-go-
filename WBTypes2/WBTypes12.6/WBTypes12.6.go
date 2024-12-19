package main

import "fmt"

func main() {
	//data := map[int]int{1: 10, 2: 20, 3: 30}
	//
	//for key, value := range data {
	//	value = 100
	//	fmt.Printf("Key: %d, Local Value: %d\n", key, value)
	//}
	//
	//fmt.Println("Original Map:", data)

	data := map[int]int{1: 10, 2: 20, 3: 30}

	for key, value := range data {
		data[key] = value * 2
	}

	fmt.Println("Updated Map:", data)

}
