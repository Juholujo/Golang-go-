package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Определяем интерфейс Runner
type Runner interface {
	Run()
}

// Пример структуры, реализующей интерфейс Runner
type Dog struct {
	Name string
}

// Реализация метода Run для структуры Dog
func (d Dog) Run() {
	fmt.Println(d.Name, "is running")
}

func inspectInterface(i Runner) {
	// Используем reflect для изучения интерфейса
	rv := reflect.ValueOf(i)
	rt := reflect.TypeOf(i)

	// Печатаем информацию об интерфейсе
	fmt.Println("=== Interface Inspection ===")
	fmt.Printf("Type: %v\n", rt)  // Тип интерфейса
	fmt.Printf("Value: %v\n", rv) // Значение интерфейса
	//fmt.Printf("Is nil: %v\n", rv.IsNil())    // Проверка на nil через reflect
	fmt.Printf("Underlying value: %+v\n", rv) // Значение конкретного объекта
}

func inspectUnsafeInterface(i Runner) {
	// Получаем доступ к низкоуровневому представлению интерфейса через unsafe.Pointer
	type iface struct {
		tab  unsafe.Pointer // Указатель на таблицу методов (тип интерфейса)
		data unsafe.Pointer // Указатель на конкретное значение
	}

	ptr := (*iface)(unsafe.Pointer(&i))

	fmt.Println("=== Unsafe Interface Inspection ===")
	fmt.Printf("Type pointer (itab): %v\n", ptr.tab)   // Указатель на тип
	fmt.Printf("Value pointer (data): %v\n", ptr.data) // Указатель на значение
}

func main() {
	// Вариант 1: Runner is nil
	var runner Runner
	inspectInterface(runner)
	inspectUnsafeInterface(runner)

	// Присваиваем объект Dog интерфейсу
	runner = Dog{Name: "Buddy"}
	fmt.Println("\nAfter assigning Dog{Name: \"Buddy\"}:")
	inspectInterface(runner)
	inspectUnsafeInterface(runner)

	// Проверка на nil
	if runner == nil {
		fmt.Println("Runner is nil")
	} else {
		fmt.Printf("Runner has value %v\n", runner)
	}
}
