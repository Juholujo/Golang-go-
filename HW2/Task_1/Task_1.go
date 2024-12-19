// Создайте программу для управления библиотекой
// Каждый экземпляр книги должен иметь название, автора, год издания и статус (доступна или на руках у читателя)
//Добавьте возможность добавления новой книги, поиска книги по названию, выдачи книги читателю и возврата книги
package main

import (
    "fmt"
)

type Book struct {
    Title     string
    Author    string
    Year      int
    Available bool
}

type Library struct {
    Books []Book
}

// Добавить книгу
func (l *Library) Add(title, author string, year int) {
    newBook := Book{
        Title:     title,
        Author:    author,
        Year:      year,
        Available: true,
    }
    l.Books = append(l.Books, newBook)
}

// Найти книгу по названию (возвращает индекс или -1)
func (l *Library) FindBookByTitle(title string) int {
    for i, b := range l.Books {
        if b.Title == title {
            return i
        }
    }
    return -1
}

// Выдать книгу читателю
func (l *Library) Issue(title string) bool {
    idx := l.FindBookByTitle(title)
    if idx == -1 {
        return false
    }
    if l.Books[idx].Available {
        l.Books[idx].Available = false
        return true
    }
    return false
}

// Вернуть книгу
func (l *Library) Return(title string) bool {
    idx := l.FindBookByTitle(title)
    if idx == -1 {
        return false
    }
    if !l.Books[idx].Available {
        l.Books[idx].Available = true
        return true
    }
    return false
}

// Вывести все книги
func (l *Library) PrintAllBooks() {
    for _, b := range l.Books {
        status := "Доступна"
        if !b.Available {
            status = "На руках у читателя"
        }
        fmt.Printf("Название: %s, Автор: %s, Год: %d, Статус: %s\n",
            b.Title, b.Author, b.Year, status)
    }
}

func main() {
    var lib Library

    // Добавим несколько книг
    lib.Add("Гарри Поттер и философский камень", "Дж. К. Роулинг", 1997)
    lib.Add("Война и мир", "Лев Толстой", 1869)
    lib.Add("Преступление и наказание", "Федор Достоевский", 1866)

    // Попытка выдачи книги
    issued := lib.Issue("Война и мир")
    if issued {
        fmt.Println("Книга 'Война и мир' выдана читателю")
    } else {
        fmt.Println("Не удалось выдать книгу 'Война и мир'")
    }

    // Попытка вернуть книгу
    returned := lib.Return("Война и мир")
    if returned {
        fmt.Println("Книга 'Война и мир' возвращена")
    } else {
        fmt.Println("Не удалось вернуть книгу 'Война и мир'")
    }

    // Вывод всех книг
    fmt.Println("Список всех книг:")
    lib.PrintAllBooks()
}
