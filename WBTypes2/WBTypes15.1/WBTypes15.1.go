package main

import "fmt"

// Функция, которая складывает два числа и возвращает результат
func sum(a, b int) int {
	return a + b
}

// Функция, которая возвращает другую функцию
func getPrinter() func(string) {
	// Функция, которая принимает строку и печатает её
	printFunc := func(message string) {
		fmt.Println("Вложенная функция:", message)
	}
	return printFunc
}

// Функция, которая принимает число и функцию, применяет эту функцию к числу и возвращает результат
func applyFunction(num int, operation func(int) int) int {
	return operation(num)
}

// Функция, которая умножает число на два
func multiplyByTwo(num int) int {
	return num * 2
}

func main() {
	// Объявление и вызов функции, которая возвращает результат сложения двух чисел
	result := sum(2, 3)
	fmt.Println("Сумма чисел:", result)

	// Функция, которая возвращает другую функцию и затем использует её
	printer := getPrinter()
	printer("Привет из вложенной функции!")

	// Функция, которая использует другую функцию как аргумент
	doubleResult := applyFunction(5, multiplyByTwo)
	fmt.Println("Результат умножения на два:", doubleResult)
}
