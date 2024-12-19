package main

import "fmt"

func main() {
	a := make(map[string]int)

	// Проверяем ключ с использованием второго значения
	if _, exists := a["key"]; !exists {
		fmt.Println("Key does not exist") // Это корректно
	}

	// Добавляем ключ "key" с нулевым значением
	a["key"] = 0

	// Проверяем снова
	if _, exists := a["key"]; !exists {
		fmt.Println("Key does not exist") // Это не напечатается, так как ключ существует
	} else {
		fmt.Println("Key exists with value:", a["key"])
	}
}
