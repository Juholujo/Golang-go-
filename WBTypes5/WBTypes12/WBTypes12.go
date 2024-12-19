package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Открываем порт для прослушивания
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8081...")

	for {
		// Ждем подключений от клиентов
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Обрабатываем подключение
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		// Чтение данных от клиента
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected or error reading:", err)
			break
		}

		fmt.Printf("Message from client: %s", message)

		// Отправляем ответ клиенту
		_, err = conn.Write([]byte("Hello from Server\n"))
		if err != nil {
			fmt.Println("Error sending response to client:", err)
			break
		}
	}
}
