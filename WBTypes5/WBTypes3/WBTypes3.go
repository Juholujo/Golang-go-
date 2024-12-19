package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func receiveTaxi(ctx context.Context, result chan struct{}, idx int) {
	randomTime := time.Duration(rand.Intn(7000)) * time.Millisecond

	timer := time.NewTimer(randomTime)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Printf("Taxi arrived: %d\n", idx)
		result <- struct{}{}
	case <-ctx.Done():
		fmt.Printf("Taxi request canceled: %d\n", idx)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	ctx, cancel := context.WithCancel(context.Background())

	result := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			defer wg.Done()
			receiveTaxi(ctx, result, idx)
		}(i)
	}

	<-result
	cancel()

	wg.Wait()
}
