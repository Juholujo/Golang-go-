package main

import (
	"fmt"
	"time"
)

func getNumbers(c chan int) {
	for i := 0; i <= 5; i++ {
		c <- i * i * i
		time.Sleep(1000 * time.Millisecond)

	}

	close(c) // close channel
}

func main() {
	fmt.Println("Запуск main() ")
	ch := make(chan int)

	go getNumbers(ch)

	for {
		val, ok := <-ch
		fmt.Println(val, ok)
		if ok == false {
			fmt.Println(val, ok, "Выход из цикла")
			break
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("Завершена main()")
}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//var ch chan int
//
//func writer() {
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("Writer recovered from panic:", r)
//		}
//	}()
//	fmt.Println("Hello from writer")
//	for i := 2; i < 50; i += 2 {
//		ch <- i
//	}
//
//}
//
//func reader() {
//	defer func() {
//		close(ch)
//	}()
//	fmt.Println("Hello from reader")
//	var sum int
//	for val := range ch {
//		sum += val
//		fmt.Println(val, sum)
//		if sum >= 100 {
//			break
//		}
//	}
//}
//
//func main() {
//	fmt.Println("Main starts")
//	ch = make(chan int)
//	go writer()
//	go reader()
//	time.Sleep(2 * time.Second)
//	fmt.Println("Main ends")
//}
