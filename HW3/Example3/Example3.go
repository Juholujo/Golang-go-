//Пример 2 (из WBTypes15.1.go)
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(1) // Установим 1, чтобы увидеть линейный порядок вывода
    wg := sync.WaitGroup{}
    wg.Add(7)

    for i := 0; i < 9; i++ {
        i := i // Локальная копия i, чтобы корректно передать в замыкание
        go func(i int) {
            defer wg.Done()
            fmt.Println("Почему, КОЛЯ?", i)
        }(i)
    }

    wg.Wait()
    fmt.Println("Паника")
}


//Если поставить runtime.GOMAXPROCS(1), то у Go как будто только одна рука.
//Go делает всё по очереди. Даже если у тебя много горутин, они не могут реально выполняться одновременно – Go будет переключаться между ними.
//Но порядок будет всегда один и тот же, потому что всего один поток. Поэтому результат стабильный и предсказуемый. Например, всегда будет начинаться с «Почему, КОЛЯ? 8» при каждом запуске.