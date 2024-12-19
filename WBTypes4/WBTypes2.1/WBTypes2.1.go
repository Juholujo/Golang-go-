package main

import "fmt"

func main() {
	// Создаем буферизованный канал с емкостью 3
	ch := make(chan int, 3)

	// Записываем несколько значений в канал
	ch <- 10
	ch <- 20
	ch <- 30

	// Закрываем канал
	close(ch)

	// Читаем из канала
	for i := 0; i < 5; i++ {
		val, ok := <-ch
		if ok {
			fmt.Printf("Прочитано значение: %d,\n", val)
		} else {
			fmt.Println("Канал закрыт, возвращено нулевое значение", val)
		}
	}
}
