package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	fmt.Println("CPU", runtime.NumCPU())

	go func() {
		fmt.Println("fdfdf")
	}()
	//time.Sleep(1 * time.Second)
	fmt.Println("Gourutines", runtime.NumGoroutine())

}
