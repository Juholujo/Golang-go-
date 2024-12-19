package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := make([]int, 4, 7)
	array := (*[4]int)(slice)

	slice[0] = 7

	slice[3] = 7

	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	fmt.Println("VALUE slice:", reflect.ValueOf(slice))
	fmt.Println("TYPE slice:", reflect.TypeOf(slice))

	fmt.Printf("array: %v, len: %d, cap: %d\n", array, len(array), cap(array))

	fmt.Println("VALUE array:", reflect.ValueOf(array))
	fmt.Println("TYPE array:", reflect.TypeOf(array))
}

/*Создание среза: slice := make([]int, 4, 7) создает срез slice с длиной 4 и емкостью 7.
Это означает, что срез может хранить 4 элемента, но выделено место для 7.
Приведение к массиву: array := (*[3]int)(slice) пытается привести срез к указателю на массив из 3 элементов.
В данном случае, array будет указывать на первые 3 элемента среза.
Изменение значения: slice[0] = 7 устанавливает первый элемент среза в 7.
Вывод информации:
slice выводит [7 0 0 0], длина 4, емкость 7.
array выводит &[7 0 0], длина 3, емкость 3 (так как это указатель на массив).*/
