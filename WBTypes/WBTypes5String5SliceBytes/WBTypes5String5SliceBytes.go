package main

import (
	"fmt"
)

func main() {
	str := "hello" // Создание строки "hello"
	fmt.Printf("Адрес переменной str: %p\n", &str)
	// Попытка изменить символ строки по индексу
	// Это приведет к ошибке компиляции
	//str[0] = 'H'

	/*	type StringHeader struct { // 0xc00004c270 Адрес структуры
		Data uintptr // адрес памяти на h
		Len  int
	}*/

	// Попытка добавить символ к строке

	fmt.Printf("Адрес переменной str до изменения: %p\n", &str) //0xc00004c270
	str += " world"
	fmt.Println(str)
	fmt.Printf("Адрес переменной str после изменения: %p\n", &str) // 0xc00004c270

}
