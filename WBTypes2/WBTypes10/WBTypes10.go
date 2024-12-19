package main

import "fmt"

/*	Структура (structure) в программировании - это составной тип данных,
	который позволяет объединить различные типы данных в единый тип.
	Она представляет собой совокупность различных переменных (полей), каждая из которых может иметь свой тип данных.
	Структуры используются для организации и хранения связанных данных в единую сущность.
*/
// 	Определение структуры Person
type Person struct {
	Name string
	Age  int
}

func main() {
	// Создание экземпляра структуры Person
	p0 := Person{}
	fmt.Println("Значение полей пустой структуры:", p0)

	p1 := Person{"Alice", 30}
	fmt.Println("Person 1:", p1) // Вывод: Person 1: {Alice 30}

	// Доступ к полям структуры
	fmt.Println("Name:", p1.Name) // Вывод: Name: Alice
	fmt.Println("Age:", p1.Age)   // Вывод: Age: 30

	// Изменение значений полей структуры
	p1.Age = 35
	fmt.Println("Updated Person 1:", p1) // Вывод: Updated Person 1: {Alice 35}

	// Создание указателя на структуру Person
	p2 := &Person{"Bob", 25}
	fmt.Println("Person 2:", *p2) // Вывод: Person 2: {Bob 25}

	// Инициализация структуры без заполнения всех полей
	var p3 Person
	p3.Name = "Charlie"
	fmt.Println("Person 3:", p3) // Вывод: Person 3: {Charlie 0}

	// Создание структуры с использованием именованных полей
	p4 := Person{Name: "Diana", Age: 40}
	fmt.Println("Person 4:", p4) // Вывод: Person 4: {Diana 40}
}
