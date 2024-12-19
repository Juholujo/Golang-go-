package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	// Симулируем долгую работу
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done(): // Проверка, отменен ли контекст
			fmt.Println("Работа прервана:", ctx.Err()) // Печатаем причину завершения
			return
		default:
			fmt.Println("Шаг", i)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("Процесс завершен")
}

func main() {
	// Создаем контекст с таймаутом 3 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Запускаем процесс
	go process(ctx)

	// Ждем завершения работы или истечения времени
	select {
	case <-ctx.Done():
		fmt.Println("Контекст завершен:", ctx.Err()) // Печатаем ошибку завершения
	}
}
