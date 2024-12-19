package main

import "fmt"

func main() {
	// Создаем массив фиксированного размера, содержащий элементы типа interface{}
	var arr [5]interface{}

	// Заполняем массив разными типами данных
	arr[0] = 42             // int
	arr[1] = "Hello"        // string
	arr[2] = 3.14           // float64
	arr[3] = true           // bool
	arr[4] = []int{1, 2, 3} // slice

	// Выводим все элементы массива
	for i, v := range arr {
		fmt.Printf("Элемент %d: %v (Тип: %T)\n", i, v, v)
	}
}
