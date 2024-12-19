/*Теория
Горутины — это легковесные потоки (green thread) выполнения, управляемые рантаймом Go.
Горутины (goroutines) представляют параллельные операции, которые могут выполняться независимо от функции, в которой они запущены.


Запуск горутины: Используется ключевое слово go перед вызовом функции.


Синтаксис:
go functionName()
Пример:

package main

import (
"fmt"
"time"
)

// Функция, которую будем запускать как горутину
func sayHello() {
fmt.Println("Hello from goroutine")
}

func main() {
// Запуск функции как горутины
go sayHello()

}*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main function starts")

	// Запуск новой горутины
	go func() {
		fmt.Println("Goroutine starts")
		time.Sleep(time.Second)
		fmt.Println("Goroutine ends")

	}()

	//// Ожидание некоторого времени в основной горутине
	//time.Sleep(2 * time.Second)

	fmt.Println("Main function ends")

}
