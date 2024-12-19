package main

import (
	"fmt"
	"io"
	"os"
)

// logMessage записывает сообщение в writer
func logMessage(writer io.Writer, message string) {
	writer.Write([]byte(message))
}

// customLogger - структура для пользовательского логгера
type customLogger struct{}

// Реализация метода Write для customLogger
func (c customLogger) Write(p []byte) (n int, err error) {
	fmt.Print("Custom logger: ")
	return fmt.Print(string(p))
}

func main() {
	// Используем os.Stdout как io.Writer
	logMessage(os.Stdout, "This is a log message to stdout\n")

	// Создаем файл и используем его как io.Writer
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	logMessage(file, "This is a log message to a file\n")

	// Используем customLogger как io.Writer
	var logger io.Writer = customLogger{}
	logMessage(logger, "This is a log message to custom logger\n")

}
