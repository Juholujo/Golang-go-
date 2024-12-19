package main

import (
	"context"
	"fmt"
)

func makeRequest(ctx context.Context) {
	oldValue, ok := ctx.Value("ID").(string)
	if ok {
		fmt.Println(oldValue)
	}

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	newValue, ok := newCtx.Value("ID").(string)
	if ok {
		fmt.Println(newValue)
	}
}

func main() {
	traceCtx := context.WithValue(context.Background(), "ID", "ANDREW")
	makeRequest(traceCtx)
}
