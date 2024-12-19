package main

import (
	"fmt"
	"time"
)

// Функция bar вызывает панику
func bar() {
	fmt.Println("Внутри функции bar")
	panic("Паника в функции bar") // Вызов panic

}

// Функция foo вызывает функцию bar в цикле
func foo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановлено в foo:", r)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Итерация", i)
		go bar() // Вызов функции bar в отдельной горутине
		//time.Sleep(1 * time.Second)

	}

	// Ожидание завершения горутин
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println("Начало функции main")
	foo()
	fmt.Println("Конец функции main")
}
