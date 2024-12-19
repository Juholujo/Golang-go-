//Пример 1 (из WBTypes15.2.go)
package main

import (
    "fmt"
)

func main() {
    var numbers []*int
    for _, value := range []int{10, 20, 30, 40} {
        numbers = append(numbers, &value)
    }

    for _, number := range numbers {
        fmt.Printf("%d ", *number)
    }
}


//В цикле берётся адрес одной и той же переменной value. Получается, на каждой итерации value просто меняет значение, но адрес остаётся прежним.
//Поэтому, когда проходимся по массиву указателей, все они указывают на одно и то же место в памяти, где в итоге лежит последнее значение (40).
//Если нужно, чтобы каждый элемент указывал на своё значение, нужно в цикле создавать отдельную переменную.