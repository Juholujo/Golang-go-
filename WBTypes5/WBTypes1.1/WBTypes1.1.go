package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	//defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-ctx.Done():
			fmt.Println("Stopping:", ctx.Err())
			return
		}
	}
}
