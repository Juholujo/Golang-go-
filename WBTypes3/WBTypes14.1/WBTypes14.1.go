package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		runtime.Gosched() // Уступаем процессор другим горутинам
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
		runtime.Gosched() // Уступаем процессор другим горутинам
	}
}

func main() {
	go printNumbers() // Запускаем первую горутину
	go printLetters() // Запускаем вторую горутину

	// Даём время горутинам завершиться
	time.Sleep(1 * time.Second)
}
