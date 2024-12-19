package main

import (
	"fmt"
)

func main() {
	fmt.Println("Начало функции main")

	// Отложенная функция для восстановления после паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановлено после паники:", r)
		}
	}()

	fmt.Println("Готовимся вызвать панику")
	panic("Что-то пошло не так")
	fmt.Println("Эта строка не будет выполнена")

	fmt.Println("Конец функции main")
}
