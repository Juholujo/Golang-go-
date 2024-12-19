// Необходимо найти минимальное и максимальное чётные значения, делящиеся на 2 без остатка
package main

import "fmt"

func main() {
    numbers := []int{8, 44, 3, 5, 11, 8, 2, 10, 6, 77, 15, 12}

    var minEven, maxEven *int
    for _, v := range numbers {
        if v%2 == 0 {
            if minEven == nil || v < *minEven {
                minEven = &v
            }
            if maxEven == nil || v > *maxEven {
                maxEven = &v
            }
        }
    }

    if minEven != nil && maxEven != nil {
        fmt.Println("Минимальное чётное число:", *minEven)
        fmt.Println("Максимальное чётное число:", *maxEven)
    } else {
        fmt.Println("Нет чётных чисел в слайсе.")
    }
}
