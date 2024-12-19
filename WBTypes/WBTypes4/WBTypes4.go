package main

import "fmt"

func main() {

	var MyNewInt int8
	//var MyNewInt int16
	MyNewInt = 127

	fmt.Println(MyNewInt)

	//MyNewInt = 129 // '129' (type untyped int) cannot be represented by the type int8
	//
	//fmt.Println(MyNewInt)

	/*	Сообщение об ошибке указывает на то, что значение, которое вы пытаетесь присвоить переменной типа int8,
		на самом деле превышает максимальное значение для этого типа.
	*/

}
