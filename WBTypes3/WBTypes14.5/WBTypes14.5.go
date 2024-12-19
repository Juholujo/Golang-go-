package main

import "fmt"

func somePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановлено внутри somePanic:", r)
		}
	}()
	panic("Паника в somePanic()")
	fmt.Println("Эта строка не будет выполнена в somePanic")
}

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Восстановлено внутри main:", r)
	//	}
	//}()

	fmt.Println("Начало функции main")

	somePanic() // Вызов функции, вызывающей панику

	recError := recover()                  // Восстановление после паники не сработает здесь
	fmt.Println("Значение", recError)      // Эта строка не будет выполнена, если паника не была перехвачена выше
	fmt.Println("Завершение функции main") // Эта строка не будет выполнена
}
