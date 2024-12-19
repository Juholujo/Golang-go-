package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Println("Starting server on http://localhost:8000...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик для пути "/" - увеличивает count и выводит URL.Path
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request on /")
	count++
	log.Printf("Handler '/' called. Current count: %d\n", count) // Логируем текущее значение count
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Обработчик для пути "/count" - возвращает количество запросов
func counter(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request on /count")
	log.Printf("Handler '/count' called. Returning count: %d\n", count) // Логируем вызов обработчика /count
	fmt.Fprintf(w, "Count = %d\n", count)
}
