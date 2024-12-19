package main

import "fmt"

// Функция appendPrefix добавляет префикс "LEX101" к переданной строке и возвращает неименнованный результат.
func appendPrefix(text string) string {

	return "LEX101 " + text
}

// Функция appendPrefixWithQuestion добавляет префикс "Is it" к переданной строке и возвращает результат и nil в качестве ошибки.
func appendPrefixWithQuestion(text string) (string, error) {
	return "Is it " + text + "?", nil
}

/*
Функция appendPrefixWithLength добавляет префикс "LEX101" к переданной строке и возвращает результат, а также длину этого результата.

Для именнованных значений в GO сразу определяются дефолтые значения типа, которые определеные в результате
*/
func appendPrefixWithLength(text string) (output string, length int) {
	// Печатаем значения переменных после их объявления.
	fmt.Println("output:", output) // Вывод: output: ""
	fmt.Println("length:", length) // Вывод: length: 0

	output = "LEX101" + text
	length = len(output)
	return
}

func main() {
	// Пример использования функции appendPrefix.
	modifiedText := appendPrefix("is my friend")
	fmt.Println(modifiedText) // Вывод: LEX101 is my friend

	// Пример использования функции appendPrefixWithQuestion.
	modifiedText, err := appendPrefixWithQuestion("error")
	fmt.Println(err)          // Вывод: nil, так как ошибки нет
	fmt.Println(modifiedText) // Вывод: Is it error?

	result, length := appendPrefixWithLength(" is my friend")

	// Выводим на экран значения возвращаемых переменных.
	fmt.Println("Result:", result) // Выводит: Result: LEX101 is my friend
	fmt.Println("Length:", length) // Выводит: Length: 17
}
