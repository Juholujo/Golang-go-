//Программа, которая поможет пользователям учитывать свои ежемесячные расходы
//Используйте карту, где ключом будет название категории расходов (например, "Продукты", "Транспорт", "Развлечения"), а значением - сумма расходов по этой категории
package main

import "fmt"

type Expenses struct {
    data map[string]float64
}

func NewExpenses() *Expenses {
    return &Expenses{data: make(map[string]float64)}
}

// Добавление расходов в категорию
func (e *Expenses) addExpense(category string, amount float64) {
    e.data[category] += amount
}

// Подсчет общей суммы
func (e *Expenses) total() float64 {
    var sum float64
    for _, val := range e.data {
        sum += val
    }
    return sum
}

func main() {
    expenses := NewExpenses()

    expenses.addExpense("Продукты", 200.50)
    expenses.addExpense("Транспорт", 100.00)
    expenses.addExpense("Развлечения", 350.75)
    // Добавим еще расходов в ту же категорию "Продукты"
    expenses.addExpense("Продукты", 50.25)

    fmt.Println("Общие расходы за месяц:", expenses.total())
    fmt.Println("Детализация по категориям:")
    for cat, val := range expenses.data {
        fmt.Printf("%s: %.2f\n", cat, val)
    }
}
