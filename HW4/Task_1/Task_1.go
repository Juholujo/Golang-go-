// Напишите программу, которая принимает строки через канал и подсчитывает количество слов в каждой строке.
// Используйте несколько горутин для обработки строк и один канал для передачи результатов.
package main

import (
    "fmt"
    "strings"
    "sync"
)

func main() {
    // Канал для передачи строк
    linesCh := make(chan string)
    // Канал для результатов
    resultsCh := make(chan int)

    var wg sync.WaitGroup

    // Запустим несколько горутин-обработчиков (например, 3)
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for line := range linesCh {
                words := strings.Fields(line)
                resultsCh <- len(words)
            }
        }()
    }

    // В отдельной горутине закроем resultsCh, когда обработчики закончат
    go func() {
        wg.Wait()
        close(resultsCh)
    }()

    // Отправляем строки
    inputLines := []string{
        "Всем привет!",
        "Следующая лекция в среду",
        "Увидимся на лекции!",
    }

    go func() {
        for _, l := range inputLines {
            linesCh <- l
        }
        close(linesCh) // Закрываем канал, чтобы обработчики завершили работу
    }()

    // Читаем результаты
    for count := range resultsCh {
        fmt.Printf("Word count: %d\n", count)
    }
}


// Используем один канал для строк, и несколько горутин читают из него и пишут результат в другой канал.
