package main

import (
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func biggestRespectToKolya() {
	// Подключение к базе данных
	log.Println("Connecting to DB...")

	{
		start := time.Now()
		defer timeTrack(start, "first query")
		// Первый запрос к базе данных
		time.Sleep(1 * time.Second) // Симуляция выполнения запроса
		log.Println("First query executed")
	}

	{
		start := time.Now()
		defer timeTrack(start, "second query")
		// Второй запрос к базе данных
		time.Sleep(2 * time.Second) // Симуляция выполнения запроса
		log.Println("Second query executed")
	}
}

func main() {
	biggestRespectToKolya()
}
