package main

import (
	"fmt"
	"reflect"
)

func goWB(a **int, p *int) {
	*a = p // Меняем указатель, на который ссылается `a`
	//  fdx43412

	fmt.Println("VALUE A WB:", reflect.ValueOf(a))

	fmt.Println("TYPE A WB:", reflect.TypeOf(a))

	fmt.Println("VALUE P WB:", reflect.ValueOf(p))

	fmt.Println("TYPE P goWB:", reflect.TypeOf(p))
}

func main() {
	a := 5
	p := &a // Указатель на `a`
	b := 42

	fmt.Println("VALUE a:", reflect.ValueOf(a))

	fmt.Println("TYPE a:", reflect.TypeOf(a))

	fmt.Println("VALUE p:", reflect.ValueOf(p))

	fmt.Println("TYPE p:", reflect.TypeOf(p))

	goWB(&p, &b) // Передаем адрес `p` (указателя на `a`) и адрес `b`
	//fmt.Println("KOLYA", &p == &b)
	fmt.Println(*p) // Разыменовываем `p`, чтобы получить новое значение, на которое он указывает
}
