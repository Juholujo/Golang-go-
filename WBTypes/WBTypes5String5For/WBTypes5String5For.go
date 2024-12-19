package main

import "fmt"

func main() {
	str := "Привет, Юрчик"

	// Итерация по байтам
	fmt.Println("Итерация по байтам:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("Байт с индексом %d: %v\n", i, str[i])
	}

	// Итерация по рунам
	fmt.Println("\nИтерация по рунам:")
	for i, r := range str {
		fmt.Printf("Символ с индексом %d: %c, Код: %d\n", i, r, r)
	}
}
