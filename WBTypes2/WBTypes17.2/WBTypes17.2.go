package main

import "fmt"

/*Cтруктуры могут содержать другие структуры, методы могут работать с этими вложенными структурами.*/

type Address struct {
	City, State string
}

// Person структура, которая содержит Address
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Метод для Person, который выводит полное имя и адрес
func (p Person) displayInfo() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	fmt.Printf("Address: %s, %s\n", p.Address.City, p.Address.State)
}

// Метод для Person, который изменяет адрес
func (p *Person) updateAddress(city, state string) {
	p.Address.City = city
	p.Address.State = state
}

func main() {
	// Создаем экземпляр Person
	p := Person{Name: "Nickolay", Age: 30, Address: Address{City: "New York", State: "NY"}}

	// Выводим начальную информацию
	p.displayInfo()

	// Обновляем адрес
	p.updateAddress("San Francisco", "CA")

	// Снова выводим информацию, чтобы увидеть изменения
	p.displayInfo()
}
