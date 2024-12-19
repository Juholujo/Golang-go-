package main

import "fmt"

func main() {
	a := make(map[string]int)

	// Ключа "key" нет, выводим сообщение
	if a["key"] == 0 {
		fmt.Println("Key does not exist") // Это корректный вывод
	}

	// Добавляем ключ "key" с нулевым значением
	a["key"] = 0

	// Снова проверяем условие
	if a["key"] == 0 {
		fmt.Println("Key does not exist") // Ошибка! Ключ существует, но его значение равно 0
	}
}
