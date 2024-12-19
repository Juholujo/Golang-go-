package main

import "fmt"

func main() {
	// Создаем set на основе мапы
	set := make(map[int]struct{})

	// Добавляем элементы
	set[1] = struct{}{}
	set[2] = struct{}{}
	set[3] = struct{}{}

	fmt.Println("set:", set)
	// Проверяем наличие элемента
	_, exists := set[2]
	fmt.Println("2 exists in set:", exists) // true

	// Удаляем элемент
	delete(set, 2)

	// Проверяем снова
	_, exists = set[2]
	fmt.Println("2 exists in set after deletion:", exists) // false
}
