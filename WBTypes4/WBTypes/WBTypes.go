/*КАНАЛЫ
Каналы в Go — это мощный механизм для обмена данными между разными горутинами.

Каналы создаются с использованием встроенной функции make, указывая тип данных, который будет передаваться через канал.


// Создание канала для передачи целых чисел
ch := make(chan int)

// Создание канала для передачи строк
chStr := make(chan string)
Основные операции с каналами:
Отправка данных в канал (<-):

Для отправки данных в канал используется оператор <-.
ch <- 42  // Отправка значения 42 в канал ch (для канала типа int)
Получение данных из канала (<-):

Для получения данных из канала также используется оператор <-.

val := <-ch  // Получение значения из канала ch и сохранение его в переменной val (для канала типа int)
Закрытие канала:

Канал можно закрыть, чтобы сообщить получающей стороне, что больше данных не будет.

close(ch)*/

package main

import (
	"fmt"
)

func main() {
	showNilChannel()

}

func showNilChannel() {
	var nilChannel chan int
	fmt.Println("Значение nilChannel ", nilChannel)
	fmt.Printf("Len: %d Cap: %d\n", len(nilChannel), cap(nilChannel))

	// Запись в nil channel
	//nilChannel <- 1

	// Чтение из nil channel
	//<-nilChannel

	// Закрытие nil channel
	//close(nilChannel)
	ch := make(chan int)
	fmt.Println("Значение nilChannel", ch)
	fmt.Printf("Len: %d Cap: %d\n", len(ch), cap(ch))

	// Закрытие nil channel
	close(ch)
}
