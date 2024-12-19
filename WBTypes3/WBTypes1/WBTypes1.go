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

// Функция изменяет структуру, переданную по указателю
func updateByPointer(p *Person) {
	fmt.Println("Inside updateByPointer - before change:")
	printDetails(p)

	p.Name = "John" // Изменяем оригинальную структуру
	p.Age = 30

	fmt.Println("Inside updateByPointer - after change:")
	printDetails(p)
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

	fmt.Println("Before updateByPointer:")
	printDetails(&person)

	// Передаем структуру по указателю
	updateByPointer(&person)

	fmt.Println("After updateByPointer:")
	printDetails(&person) // Проверяем изменения в оригинале
}
