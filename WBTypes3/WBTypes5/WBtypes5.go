package main

import (
	"fmt"
)

// Определяем интерфейс Writer
type Writer interface {
	Write(p []byte) (n int, err error)
}

// Определяем интерфейс Closer
type Closer interface {
	Close() error
}

// Определяем интерфейс ReadWriteCloser, включая Writer и Closer
type ReadWriteCloser interface {
	Writer
	Closer
	Read(p []byte) (n int, err error)
}

// Пример реализации структуры File, которая реализует интерфейс ReadWriteCloser
type File struct{}

func (f File) Write(p []byte) (n int, err error) {
	// Реализация метода Write для File
	fmt.Println("Writing:", string(p))
	return len(p), nil
}

func (f File) Close() error {
	// Реализация метода Close для File
	fmt.Println("Closing the file")
	return nil
}

func (f File) Read(p []byte) (n int, err error) {
	// Реализация метода Read для File
	fmt.Println("Reading into buffer")
	return len(p), nil
}

func main() {
	// Создаем экземпляр File
	file := File{}

	// Пример использования интерфейса Writer
	data := []byte("Hello, World!")
	n, err := file.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Bytes written: %d\n", n)

	// Пример использования интерфейса Closer
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing:", err)
		return
	}

	// Пример использования интерфейса ReadWriteCloser
	var rw ReadWriteCloser = file // Присваиваем переменной тип ReadWriteCloser экземпляр File
	buffer := make([]byte, 1024)
	n, err = rw.Read(buffer) // Используем метод Read интерфейса ReadWriteCloser
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Bytes read: %d\n", n)
}
