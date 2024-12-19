package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User представляет данные пользователя
type User struct {
	ID    int    `json:"idff"`
	Name  string `json:"nameff"`
	Email string `json:"email"`
}

func main() {
	// Пример структуры для сериализации
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Сериализация структуры в JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	fmt.Println("Serialized JSON:", string(jsonData))

	// Пример JSON строки для десериализации
	jsonStr := `{"idff":1,"nameff":"John Doe","email":"john.doe@example.com"}`

	// Десериализация JSON в структуру
	var userFromJSON User
	if err := json.Unmarshal([]byte(jsonStr), &userFromJSON); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	fmt.Printf("Deserialized User: %+v\n", userFromJSON)
}
