package main

import (
	"context"
	"fmt"
)

func main() {
	// Создаем контекст с пользовательским значением
	ctx := context.WithValue(context.Background(), "user", "Kapitalni lepi Kolja, hvala")

	// Печатаем значение в контексте
	fmt.Println("Original value:", ctx.Value("user"))

	// Попытка изменить значение (это невозможно, контекст неизменяем)
	//ctx.Value("user") = "Alice" // Ошибка: нельзя изменить значение контекста напрямую

	// Чтобы изменить значение, создаем новый контекст с новым значением
	newCtx := context.WithValue(ctx, "user", "Alice")

	// Печатаем значения из обоих контекстов
	fmt.Println("Original context value:", ctx.Value("user"))
	fmt.Println("New context value:", newCtx.Value("user"))
}
