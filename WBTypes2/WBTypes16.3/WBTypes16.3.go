package main

import (
	"fmt"
	"time"
)

func tryTest() func() {
	fmt.Println("tryTest")
	return func() {
		fmt.Println("tryTest2")
	}
}

func main() {
	defer fmt.Println("Первое время:", time.Now())
	defer tryTest()()
	time.Sleep(2 * time.Second)
	defer fmt.Println("Второе время", time.Now())

}
