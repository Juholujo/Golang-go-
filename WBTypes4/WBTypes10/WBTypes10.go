package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	mu      sync.Mutex // Объявляем мьютекс
)

func increment() {
	mu.Lock()         // Захватываем мьютекс перед доступом к общим данным
	defer mu.Unlock() // Освобождаем мьютекс в конце функции

	counter++ // Увеличиваем значение счетчика на 1
}

func main() {
	// Запускаем несколько горутин, которые будут инкрементировать счетчик
	for i := 0; i < 10000000; i++ {
		go increment()
	}

	// Даем времени горутинам выполниться
	time.Sleep(time.Second)

	fmt.Println("Итоговое значение счетчика:", counter)
}
