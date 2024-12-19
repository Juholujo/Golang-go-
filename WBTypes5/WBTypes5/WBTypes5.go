package main

import (
	"context"
	"fmt"
)

func main() {
	traceCtx := context.WithValue(context.Background(), "trace_id", "18-12-2024")
	makeRequest(traceCtx)

	oldValue, ok := traceCtx.Value("trace_id").(string)
	if ok {
		fmt.Println("mainValue", oldValue)
	}
}

func makeRequest(ctx context.Context) {
	oldValue, ok := ctx.Value("trace_id").(string)
	if ok {
		fmt.Println("oldValue", oldValue)
	}

	newCtx := context.WithValue(ctx, "trace_id", "11-11-11")
	newValue, ok := newCtx.Value("trace_id").(string)
	if ok {
		fmt.Println("newValue", newValue)
	}
}
