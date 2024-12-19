package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*type StringHeader struct {
	Data uintptr
	Len  int
}*/

func main() {
	str := "hello"
	fmt.Printf("Адрес переменной str: %p\n", &str)

	// Преобразуем строку в StringHeader
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("Адрес строки (Data): 0x%x\n", header.Data)

	// Адрес первого символа строки (str[0])
	firstCharAddress := header.Data
	fmt.Printf("Адрес первого символа str[0]: 0x%x\n", firstCharAddress)

	// Для проверки: конвертируем адрес обратно в указатель и читаем значение
	firstChar := *(*byte)(unsafe.Pointer(firstCharAddress))
	fmt.Printf("Первый символ str[0]: %c\n", firstChar)
}
