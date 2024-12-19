package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 9; i++ {
		i := i
		go func(i int) {
			defer wg.Done()
			fmt.Println("                           Почему, КОЛЯ?", i)
		}(i)
	}
	wg.Wait()

	// time.Sleep(1 * time.Second)
	fmt.Println("Паника")
}
