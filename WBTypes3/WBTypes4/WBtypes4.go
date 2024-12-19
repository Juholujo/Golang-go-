/*Интерфейс в языке программирования Go представляет собой тип данных,
описывает набор методов, которые должен реализовать тип,
чтобы считаться реализующим этот интерфейс. (представляет из себя набор сигнатур методов,
которые нужно реализовать его имплементации)
Имплементация интерфейса в  означает создание конкретной реализации всех методов, определенных в интерфейсе.




type Writer interface {
Write(p []byte) (n int, err error)
}
В приведенном примере Writer - это интерфейс,
у него один метод Write, который принимает слайс байтов и возвращает количество записанных байтов и ошибку.

Реализация интерфейса:

Чтобы тип данных в Go считался реализующим интерфейс, он должен реализовать все методы, определенные в этом интерфейсе.
Это происходит автоматически: если тип имеет все необходимые методы, он автоматически считается реализующим интерфейс.


type myWriter struct{}

func (mw myWriter) Write(p []byte) (n int, err error) {
// Реализация метода Write
return len(p), nil
}
В этом примере myWriter реализует метод Write, определенный в интерфейсе Writer.*/

package main

import "fmt"

// Определяем интерфейс Animal
type Animal interface {
	Speak() string // Интерфейс требует реализации метода Speak, который возвращает строку
}

// Тип Dog
type Dog struct {
	Name string
}

// Реализация метода Speak для типа Dog
func (d Dog) Speak() string {
	return "Woof!" // Собака говорит "Woof!"
}

// Тип Cat
type Cat struct {
	Name string
}

// Реализация метода Speak для типа Cat
func (c Cat) Speak() string {
	return "Meow!" // Кошка говорит "Meow!"
}

func main() {
	var animals []Animal

	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	animals = append(animals, dog, cat) // Добавляем собаку и кошку в срез интерфейсов Animal

	for _, animal := range animals {
		fmt.Println(animal.Speak()) // Вызываем метод Speak для каждого животного
	}

	fmt.Println(dog.Speak()) // Вывод: Woof!
	fmt.Println(cat.Speak()) // Вывод: Meow!
	fmt.Println(myInterface)
	myInterface.SomeMethod() // // Вызов метода на nil интерфейсе вызовет панику
}

type MyInterface interface {
	SomeMethod()
}

var myInterface MyInterface

// Вызов метода на nil интерфейсе вызовет панику

//Когда интерфейс равен nil, вызовы его методов приводят к панике (runtime panic).
//Это происходит потому, что интерфейс не имеет значения (в том числе и значения, которое представляет собой получателя метода),
//поэтому невозможно выполнить вызов метода на nil интерфейсе.
