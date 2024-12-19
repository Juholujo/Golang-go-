// Нужно объяснить вывод кода с defer
package main

import (
    "fmt"
    "time"
)

func tryTest() func() {
    fmt.Println("tryTest")
    return func() {
        fmt.Println("tryTest2")
    }
}

func main() {
    defer fmt.Println("Первое время:", time.Now())
    defer tryTest()()
    time.Sleep(2 * time.Second)
    defer fmt.Println("Второе время", time.Now())
}


// При defer tryTest():
// Если просто написать defer tryTest() без ():
// При постановке на defer вызывается tryTest(), печатается "tryTest"
// Возвращенная функция не будет вызвана, так как мы не поставили дополнительные ()
// В конце при выходе из main не будет "tryTest2", т.к. ее вызова не было. Просто отложится сам факт вызова tryTest()
// который уже произошел при добавлении defer, но поскольку возвращенная функция не вызвана — "tryTest2" не выводится

// При defer tryTest()():
// Когда выполняется строка defer tryTest()(), сначала вызывается tryTest()
// tryTest() выводит "tryTest" немедленно и возвращает анонимную функцию, которая выводит "tryTest2"
// Затем () после tryTest() означает, что мы сразу же вызываем возвращенную функцию,
// но так как это стоит после ключевого слова defer, вызов этой возвращенной функции (tryTest2) откладывается до конца main

