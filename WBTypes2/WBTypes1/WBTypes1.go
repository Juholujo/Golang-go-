package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Первый пример:

	list := make([]int, 4, 5)
	fmt.Println("Before append:", list)

	//fmt.Println("Before append cap:", list[:cap(list)])
	//list = list[:cap(list)]
	//list[4] = 11

	//fmt.Println("After append 11:", list)

	// Получаем адрес массива данных `list` через SliceHeader
	headerList := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	fmt.Printf("Memory list: Data=0x%x Len=%d Cap=%d\n", headerList.Data, headerList.Len, headerList.Cap)

	// Добавляем элемент в `list`, возвращаемый срез - `list2`
	list2 := append(list, 1)

	fmt.Println("after append 1:", list2)
	fmt.Println("list 1:", list)

	fmt.Println("after append 1 list:", list[:cap(list)])

	headerList2 := (*reflect.SliceHeader)(unsafe.Pointer(&list2))
	fmt.Printf("Memory list2: Data=0x%x Len=%d Cap=%d\n", headerList2.Data, headerList2.Len, headerList2.Cap)

	// Изменяем значение в исходном срезе
	list[0] = 4
	fmt.Println("After list[0] = 4:")
	fmt.Println("list:", list)
	fmt.Println("list2:", list2)

	// Изменяем значение в `list2`
	list2[0] = 7
	fmt.Println("After list2[0] = 7:")
	fmt.Println("list:", list)
	fmt.Println("list2:", list2)
}
