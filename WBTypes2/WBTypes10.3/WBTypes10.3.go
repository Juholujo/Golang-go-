package main

import (
	"fmt"
)

// Определение структуры Persons
type Persons struct {
	Name  string
	Age   int
	Email string
	Phone string
	Male  bool
}

// Определение структуры NewPerson с полем Persons типа []Persons
type NewPersons struct {
	Persons []Persons
}

func main() {
	// Создание новой структуры NewPersons

	/*Поле Persons в структуре NewPersons имеет тип []Persons, что означает слайс  структур типа Persons.
	Таким образом, мы инициализируем это поле с помощью литерала слайса []Persons{}, внутри которого перечислены экземпляры структуры Persons.

	Каждый экземпляр структуры Persons внутри литерала слайса также инициализируется с помощью фигурных скобок {},
	где указываются значения для каждого поля этой структуры */

	np := NewPersons{
		Persons: []Persons{
			{
				Name:  "Yurchik",
				Age:   30,
				Email: "john@example.com",
				Phone: "123-456-7890",
				Male:  true,
			},
			{
				Name:  "David",
				Age:   25,
				Email: "jane@example.com",
				Phone: "987-654-3210",
				Male:  false,
			},
		},
	}

	// Вывод значений полей структур Persons внутри среза Persons
	for _, p := range np.Persons {
		fmt.Println("Name:", p.Name)
		fmt.Println("Age:", p.Age)
		fmt.Println("Email:", p.Email)
		fmt.Println("Phone:", p.Phone)
		fmt.Println("Male:", p.Male)
	}

	fmt.Println("Вывод значений полей структур Persons через np.Persons с использованием индексации")
	// Вывод значений полей структур Persons через np.Persons с использованием индексации
	for i := 0; i < len(np.Persons); i++ {
		fmt.Println("Name:", np.Persons[i].Name)
		fmt.Println("Age:", np.Persons[i].Age)
		fmt.Println("Email:", np.Persons[i].Email)
		fmt.Println("Phone:", np.Persons[i].Phone)
		fmt.Println("Male:", np.Persons[i].Male)
	}
}
