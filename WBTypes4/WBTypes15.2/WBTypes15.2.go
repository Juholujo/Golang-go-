package main

import (
	"context"
	"fmt"
	"time"
)

func longTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Проверка, отменен ли контекст
			fmt.Println("Задача отменена:", ctx.Err()) // Печатаем причину завершения
			return
		default:
			// Долгая работа
			fmt.Println("Работаю...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	// Создаем контекст, который можно отменить вручную
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем долгую задачу
	go longTask(ctx)

	// Ждем 5 секунд и затем отменяем контекст
	time.Sleep(5 * time.Second)
	cancel() // Отменяем контекст

	// Ждем завершения работы горутины
	time.Sleep(1 * time.Second)
}
