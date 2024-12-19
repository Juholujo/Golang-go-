package main

import "fmt"

func main() {
	/*Мы можем получить подстроку*/

	str := "Hello!"

	// Получение подстроки, начиная с индекса 2 и заканчивая 3
	/*Начинается с from и заканчиватся to*/

	newStr := str[2:4]
	fmt.Printf("Подстрока str[2:4]: %s\n", newStr)

	// Получение подстроки, начиная с начала и заканчивая 3
	newStr = str[:4] // 0 - 3
	fmt.Printf("Подстрока str[:4]: %s\n", newStr)

	// Получение подстроки, начиная с индекса 2 и до конца
	newStr = str[2:]
	fmt.Printf("Подстрока str[2:]: %s\n", newStr)
}
