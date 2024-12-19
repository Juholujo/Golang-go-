package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

// Функция изменяет структуру, переданную по значению (копию)
func updateByValue(p Person) {
	fmt.Println("Inside updateByValue - before change:")
	printDetails(&p)

	p.Name = "John" // Изменяем копию структуры
	p.Age = 30

	fmt.Println("Inside updateByValue - after change:")
	printDetails(&p)
}

func printDetails(ptr interface{}) {
	// Отражаем значение переданного указателя
	val := reflect.ValueOf(ptr).Elem()
	typ := val.Type()

	// Выводим адрес структуры
	fmt.Printf("Address of structure: %p\n", ptr)

	// Перебираем поля структуры
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldAddr := unsafe.Pointer(field.UnsafeAddr())

		fmt.Printf("Field: %s, Value: %v, Address: %p\n", fieldType.Name, field.Interface(), fieldAddr)
	}
	fmt.Println("-----")
}

func main() {
	// Инициализируем структуру
	person := Person{Name: "Alice", Age: 25}

	fmt.Println("Before updateByValue:")
	printDetails(&person)

	// Передаем структуру по значению
	updateByValue(person)

	fmt.Println("After updateByValue:")
	printDetails(&person) // Проверяем, изменился ли оригинал
}
