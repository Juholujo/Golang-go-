package main

import "fmt"

func main() {
	ch := make(chan int) // Небуферизованный канал
	_ = ch
	go func() {
		value := <-ch // Получаем значение из канала
		fmt.Println("Received:", value)
	}()

	ch <- 42 // Отправляем значение в канал (блокируется, так как нет получателя)
	fmt.Println("Sent 42")

	//time.Sleep(1 * time.Second)
	//go func() {
	//	value2 := 52
	//	ch <- value2 // Получаем значение из канала
	//	fmt.Println("Sent 52:", value2)
	//}()
	//
	//value := <-ch // Получаем значение из канала (блокируется, так как нет отправителя)
	//fmt.Println("Received:", value)
}
