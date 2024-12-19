package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем буферизованный канал с емкостью 5
	ch := make(chan int, 5)
	fmt.Printf("Длина: %d, Емкость: %d\n", len(ch), cap(ch))

	// Функция для записи значений в канал
	go func() {
		//defer close(ch)
		for i := 0; i < 7; i++ {
			//defer func() {
			//	if r := recover(); r != nil {
			//		fmt.Println("Восстановлено внутри somePanic:", r)
			//	}
			//}()
			ch <- i
			// Выводим текущую длину и емкость канала после каждой записи
			fmt.Printf("Записано значение %d в канал. Длина: %d, Емкость: %d\n", i, len(ch), cap(ch))
			//close(ch)
		}
		//// Закрываем канал после записи всех значений
		close(ch)
	}()
	time.Sleep(2 * time.Second)
	// Чтение значений из канала
	for val := range ch {
		fmt.Printf("Прочитано значение %d из канала. Длина: %d, Емкость: %d\n", val, len(ch), cap(ch))

	}
}
