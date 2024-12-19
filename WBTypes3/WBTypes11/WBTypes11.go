package main

import (
	"fmt"
)

// Псевдоним для пустого интерфейса
type any interface{}

// Метод String реализует интерфейс Stringer для MyString
func (ms MyString) String() string {
	return ms.value
}

// Использование интерфейса Stringer
type Stringer interface {
	String() string
}

// Пример реализации интерфейса Stringer
type MyString struct {
	value string
}

func main() {
	// Пример использования пустого интерфейса
	var a any = "hello"
	fmt.Println(a) // Вывод: hello

	// Создаем мапы разных типов
	anyMap := make(map[any]any)
	intStringMap := map[int]string{
		1: "one",
		2: "two",
	}

	// Добавляем элементы в мапу anyMap
	anyMap[1] = "one"
	anyMap[2] = "two"

	// Попытка присвоить intStringMap в anyMap вызовет ошибку компиляции
	//anyMap = intStringMap // Ошибка: cannot use intStringMap (type map[int]string) as type map[any]any in assignment

	// Копируем значения из intStringMap в anyMap
	for k, v := range intStringMap {
		anyMap[k] = v
	}

	// Выводим содержимое anyMap
	fmt.Println(anyMap) // Вывод: map[1:one 2:two]

	// Создаем переменную типа Stringer и инициализируем ее экземпляром MyString
	var s Stringer = MyString{"Hello, Go!"}

	// Вызываем метод String() через интерфейс Stringer
	fmt.Println(s.String()) // Вывод: Hello, Go!

}
