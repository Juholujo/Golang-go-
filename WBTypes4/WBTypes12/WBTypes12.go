package main

import "time"

func main() {
	ch := make(chan int, 2)
	ch <- 42
	_ = <-ch
	ch <- 20
	ch <- 20
	_ = <-ch

	go func() {
		for {
			ch <- 50
		}
	}()
	time.Sleep(2 * time.Second)

	close(ch)

	return

}
