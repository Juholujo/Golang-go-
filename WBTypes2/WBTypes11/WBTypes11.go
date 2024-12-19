package main

import (
	"fmt"
)

func main() {
	/* 	Карта (map) в Go представляет собой структуру данных, которая отображает ключи на значения.
	Ключи и значения могут быть любого типа данных,
	но ключи должны быть уникальными в пределах одной карты. */

	/*	// A header for a Go map.
		type hmap struct {
			// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
			// Make sure this stays in sync with the compiler's definition.
			count     int // # live cells == size of map.  Must be first (used by len() builtin)
			flags     uint8
			B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
			noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
			hash0     uint32 // hash seed

			buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
			oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
			nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

			extra *mapextra // optional fields
		}*/

	// Объявление и инициализация карты
	var m map[string]int                     // Объявление пустой карты m, где ключи являются строками, а значения - целыми числами
	fmt.Println("Значение переменной m:", m) // Вывод: map[]

	// Инициализация карты с тремя парами ключ-значение
	m = map[string]int{"LEX101": 5, "Hainken": 3, "ALEX": 7}            // Инициализация карты с тремя парами ключ-значение
	fmt.Println("Значение переменной m после после инициализации :", m) //

	// Получение значения по ключу
	fmt.Println("Value of ALEX:", m["ALEX"]) // Вывод значения, связанного с ключом "apple": 5

	// Добавление новой пары ключ-значение
	m["grape"] = 9                                            // Добавление пары "grape": 9 в карту m
	fmt.Println("Значение переменной m после добавления:", m) // Вывод: map[apple:5 banana:3 orange:7 grape:9]

	// Удаление элемента по ключу
	delete(m, "banana")                                     // Удаление ключа "banana" из карты m
	fmt.Println("Значение переменной m после удаления:", m) // Вывод: map[apple:5 orange:7 grape:9]

	// Проверка наличия ключа в карте
	_, ok := m["apple"]                  // Проверка наличия ключа "apple" в карте m
	fmt.Println("Is apple present?", ok) // Вывод: true

	// Итерация по картам
	for key, value := range m { // Итерация по парам ключ-значение в карте m
		fmt.Println(key, ":", value) // Вывод пар ключ-значение
	}

	// Создание карты с помощью make
	m2 := make(map[int]string)                 // Создание пустой карты m2, где ключи являются целыми числами, а значения - строками
	m2[1] = "one"                              // Добавление пары ключ-значение в карту m2
	fmt.Println("Значение переменной m2:", m2) // Вывод: map[1:one]

	/*	Вложенные карты - это карты, которые являются значениями другой карты.
		Это позволяет создавать более сложные иерархические структуры данных, где каждый элемент карты может быть самой картой.
		Давайте разберем пример:
	*/
	nestedMap := map[string]map[string]int{
		"group1": {"a": 1, "b": 2},
		"group2": {"x": 10, "y": 20},
	}
	fmt.Println("Значение переменной nestedMap:", nestedMap)

	/*		Здесь nestedMap - это внешняя карта, которая содержит вложенные карты в качестве своих значений.

			Каждый ключ во внешней карте ("group1", "group2") связан с вложенной картой.
			Вложенные карты имеют ключи и значения, где ключи и значения могут быть любого типа, поддерживаемого Go.
			В данном примере, для ключа "group1" значение является вложенной картой {"a": 1, "b": 2},
			а для ключа "group2" значение - вложенная карта {"x": 10, "y": 20}.
			Таким образом, мы можем обращаться к элементам вложенных карт так же, как и к обычным картам.
			Например:*/

	fmt.Println(nestedMap["group1"]["a"]) // Вывод: 1
	fmt.Println(nestedMap["group2"]["y"]) // Вывод: 20

}
