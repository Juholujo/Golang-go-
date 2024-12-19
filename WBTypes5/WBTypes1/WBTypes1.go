package main

import (
	"fmt"
	"time"
)

func main() {
	//ticker := time.NewTicker(1 * time.Second)
	//
	//for t := range ticker.C {
	//	fmt.Println("Tick at", t)
	//
	//}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for t := range ticker.C {
			fmt.Println("Checking server status at", t)
			// Логика проверки состояния сервера.
		}
	}()

	time.Sleep(10 * time.Second) // Примерно 4 проверки.
	fmt.Println("Stopping server status check.")

}
