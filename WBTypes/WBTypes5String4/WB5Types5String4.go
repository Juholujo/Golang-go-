package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var str = "Здраво, Коля"

	fmt.Println("Len str", len(str))

	var str2 = "Zdravo Nickolay"

	fmt.Println("Len str2", len(str2))

	fmt.Println("LEN STR2", unsafe.Sizeof(len(str2)))

}

/*len(str): возвращает длину строки str в байтах.
Для строки "Здраво, Коля" результат 22, потому что в этой строке есть символы,
представленные в UTF-8, где кириллические буквы занимают 2.
len(str2): возвращает длину строки str2 в байтах.
Для строки "Zdravo Nickolay" результат 15, так как все символы занимают по 1 байту в UTF-8.
unsafe.Sizeof(len(str2)): возвращает размер переменной, которая хранит результат len(str2).
Результат 8, потому что len(str2) — это целое число (int), а на платформе  64-битной int занимает 8 байтов.*/
