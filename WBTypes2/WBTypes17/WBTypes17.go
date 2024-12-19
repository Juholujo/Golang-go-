package main

import "fmt"

/*	Методы в Go - это функции, связанные с определенным типом данных.
	Они позволяют добавить поведение к типам данных, делая их более гибкими и организованными.
	Методы позволяют работать с данными конкретного типа, вызывая функции, которые связаны с этим типом данных.
*/

/*Синтаксис метода

Метод определяется так же, как и функция, но с добавлением получателя (receiver):

func (receiverName receiverType) methodName(parameters) (results) {
	// method body
}
receiverType - тип, к которому будет привязан метод.
receiverName - имя получателя (receiver). Оно используется для доступа к полям и методам получателя внутри тела метода.
methodName - имя метода.
parameters - параметры метода (если они есть).
results - возвращаемые значения метода (если они есть)*/

// Определяем структуру User
type User struct {
	Name  string
	Email string
}

// Метод для структуры User, который выводит информацию о пользователе
// user - это получатель (receiver) метода displayInfo
func (user User) displayInfo() {
	fmt.Printf("Name: %s, Email: %s\n", user.Name, user.Email)
}

// Метод для структуры User, который изменяет Email пользователя
// Здесь используется указатель на User, чтобы изменения были видны за пределами метода
func (user *User) updateEmail(newEmail string) {
	user.Email = newEmail
}

func main() {
	// Создаем новый экземпляр структуры User
	u := User{Name: "Alice", Email: "alice@example.com"}

	// Вызываем метод displayInfo для отображения информации о пользователе
	u.displayInfo() // Вывод: Name: Alice, Email: alice@example.com

	// Вызываем метод updateEmail для изменения Email пользователя
	u.updateEmail("alice@newdomain.com")

	// Снова вызываем метод displayInfo, чтобы увидеть изменения
	u.displayInfo() // Вывод: Name: Alice, Email: alice@newdomain.com
}
