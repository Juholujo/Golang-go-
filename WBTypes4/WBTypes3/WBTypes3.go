package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch) // Закрываем канал после завершения работы горутины
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("Отправлено в канал:", i)
			time.Sleep(1 * time.Second)

		}
	}()

	for value := range ch {

		fmt.Println("Прочитано из канала:", value)

	}
}
