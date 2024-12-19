package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшает счетчик на 1, когда функция завершается
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		//wg.Add(1) // Увеличивает счетчик горутин на 1
		go worker(i, &wg)
	}
	//wg.Add(1)
	wg.Wait() // Блокирует выполнение, пока все горутины не завершатся
	fmt.Println("All workers done")
}
