package main

import "fmt"

func main() {
	/*	Также в Go отсутствует отдельный символьный тип данных, как в некоторых других языках программирования.
		Вместо этого символы представлены типом rune.
		Тип rune представляет собой 32-битное значение Unicode.
		Тип rune является псевдонимом для int32 и предназначен для использования в контексте символов Unicode.

	Пример:*/

	/*Давайте создим переменную типа rune и посмотрим на вывод*/

	var char rune = 'ᇧ'
	// https://lphp.ru/unicode_utf8/unicode.php?ysclid=lw5zlqd8tp311858194

	fmt.Println("Значение переменной char в unicode:", char)

	var char2 = 'ᇧ'

	fmt.Println("Значение переменной char в unicode:", char2)

	var char3 = "A"

	fmt.Printf("Значение переменной char3: %s, тип переменной: %T\n", char3, char3)

	var char4 = 'A'

	fmt.Printf("Значение переменной char4: %v, тип переменной: %T\n", char4, char4)

	/*	Таким образом, тип string необходим для работы с текстовыми строками в Go,
		а тип rune используется для работы с символами Unicode внутри строк.*/

	var char5 = 'ᇧ' // string(char5) - это преобразование переменной char5 в строку

	fmt.Println("Значение переменной char5:", string(char5))
}
