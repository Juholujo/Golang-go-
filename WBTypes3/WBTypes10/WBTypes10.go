package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func logMessage(writer io.Writer, message string) {
	writer.Write([]byte(message))
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

	// Используем bytes.Buffer как io.Writer
	var buffer bytes.Buffer
	logMessage(&buffer, "This is a log message to a bytes.Buffer\n")
	fmt.Println(buffer.String())

	// Используем strings.Builder как io.Writer
	var builder strings.Builder
	logMessage(&builder, "This is a log message to a strings.Builder\n")
	fmt.Println(builder.String())
}
