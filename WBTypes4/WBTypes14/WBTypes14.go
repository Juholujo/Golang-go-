package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. context.Background()
	// Это корневой пустой контекст, который используется как начальный контекст.
	ctx := context.Background()
	fmt.Println("1. context.Background():", ctx)

	// 2. context.TODO()
	// Это контекст, который используется как заглушка, если контекст еще не определен.
	ctx2 := context.TODO()
	fmt.Println("2. context.TODO():", ctx2)

	// 3. context.WithCancel()
	// Создаем контекст с возможностью отмены.
	ctx3, cancel := context.WithCancel(ctx)
	defer cancel() // Обязательно вызываем cancel, чтобы освободить ресурсы.
	fmt.Println("3. context.WithCancel():", ctx3)
	cancel() // Отменяем контекст

	// 4. context.WithTimeout()
	// Создаем контекст с таймаутом 2 секунды.
	ctx4, cancel4 := context.WithTimeout(ctx, 2*time.Second)
	defer cancel4()
	fmt.Println("4. context.WithTimeout():", ctx4)
	time.Sleep(3 * time.Second) // Ждем больше, чем таймаут, чтобы контекст отменился автоматически

	// 5. context.WithDeadline()
	// Создаем контекст с дедлайном, который сработает через 3 секунды.
	ctx5, cancel5 := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel5()
	fmt.Println("5. context.WithDeadline():", ctx5)
	time.Sleep(4 * time.Second) // Ждем больше, чем дедлайн

	// 6. context.WithValue()
	// Добавляем пользовательские данные в контекст.
	ctx6 := context.WithValue(ctx, "userID", 12345)
	fmt.Println("6. context.WithValue():", ctx6)
	// Получаем значение из контекста
	userID := ctx6.Value("userID")
	if userID != nil {
		fmt.Println("UserID из контекста:", userID)
	} else {
		fmt.Println("UserID не найден")
	}
}
