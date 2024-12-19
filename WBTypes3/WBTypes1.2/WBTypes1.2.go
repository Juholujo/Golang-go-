package main

import "fmt"

func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Привет")
	testSlice = append(testSlice, "Привет")
	test(&testSlice) // Передаем указатель на срез

	fmt.Println(testSlice) // Выводим измененный срез

}

func test(testSlice *[]string) { // Принимаем указатель на срез
	*testSlice = append(*testSlice, "WB") // Меняем содержимое по указателю
}
