package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	baseCapacity := 0
	s := make([]int, 0, baseCapacity) // Создаем срез с базовой капасити

	headerList := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("Memory list: Data=0x%x Len=%d Cap=%d\n", headerList.Data, headerList.Len, headerList.Cap)

	fmt.Printf("Начальная капасити: %d\n", cap(s))

	// Добавляем элементы и выводим капасити после каждого добавления
	for i := 0; i < 20; i++ {
		s = append(s, i)
		fmt.Printf("После добавления %d элементов, капасити: %d\n", i+1, cap(s))
	}
}
