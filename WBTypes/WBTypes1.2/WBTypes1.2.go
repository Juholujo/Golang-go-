package main

import "fmt"

type MyInt int // Создаем новый тип, основанный на int

func main() {
	var a int
	var b MyInt

	// Следующий код вызовет ошибку компиляции:
	// fmt.Println(a == b) // Ошибка компиляции: cannot compare a (variable of type int) with b (variable of type MyInt)

	// Для сравнения нужно привести типы:

	fmt.Println(a == int(b)) // Привели MyInt к типу int
}
