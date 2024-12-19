/*Boolean types*/

/*bool: Логический тип данных bool в Go представляет собой два возможных значения: true и false.
Этот тип данных используется для представления логических условий и результатов операций сравнения.
Примеры:*/

package main

import "fmt"

func main() {
	// Пример 1: Объявление логических переменных
	var b1 bool = true
	var b2 bool = false

	_, _ = b1, b2

	// Пример 2: Использование логических значений в условных операторах
	var age int = 20
	var isAdult bool = age >= 18 // true, если возраст больше или равен 18, иначе false

	if isAdult {
		fmt.Println("Человек совершеннолетний")
	} else {
		fmt.Println("Человек несовершеннолетний")
	}

	// Пример 4: Использование логических значений в логических операциях
	var isRaining bool = true
	var isSunny bool = false

	// Логическое И (AND)
	var isRainyAndSunny bool = isRaining && isSunny // false, так как идет дождь и солнце не светит
	fmt.Println("Дождь и солнце одновременно:", isRainyAndSunny)

	// Логическое ИЛИ (OR)
	var isRainyOrSunny bool = isRaining || isSunny // true, так как идет дождь или солнце светит
	fmt.Println("Дождь или солнце:", isRainyOrSunny)

	// Логическое НЕ (NOT)
	var isNotRaining bool = !isRaining // false, так как идет дождь
	fmt.Println("Не идет дождь:", isNotRaining)
}
