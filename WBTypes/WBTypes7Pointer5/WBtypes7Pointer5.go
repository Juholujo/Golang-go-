package main

import "fmt"

func main() {

	wbTech12 := "name"

	pointer3 := &wbTech12

	fmt.Printf("Тип переменной pointer3: %T\n", pointer3)

	/*  Мы можем получить доступ к значению, на которое указывает указатель.*/

	fmt.Println(*pointer3)

	fmt.Printf("Значение переменной wbTech12 до изменения: %T\n", wbTech12)

	/*  Мы можем изменить значение переменной, на которую указывает указатель.*/

	*pointer3 = "new name"

	fmt.Printf("Значение переменной wbTech12 после изменения: %T\n", wbTech12)

}
