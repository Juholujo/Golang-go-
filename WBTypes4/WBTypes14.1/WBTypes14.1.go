package main

import (
	"context"
	"fmt"
)

func main() {

	type key string
	const userKey key = "user"

	// Создаем новый контекст с значением
	ctx := context.WithValue(context.Background(), userKey, "Alice")

	// Извлекаем значение из контекста
	value := ctx.Value(userKey)
	if value != nil {
		fmt.Println("Value found in context:", value)
	} else {
		fmt.Println("No value found in context")
	}
}
