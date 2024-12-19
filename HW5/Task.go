package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
)

type Student struct {
    FullName         string `json:"fullname"`
    MathScore        int    `json:"math_score"`
    InformaticsScore int    `json:"informatics_score"`
    EnglishScore     int    `json:"english_score"`
}

var (
    admittedStudents []Student
    mu               sync.Mutex
)

func main() {
    // Маршрут для приёма заявок
    http.HandleFunc("/apply", applyHandler)

    // Маршрут для получения списка поступивших
    http.HandleFunc("/admitted", admittedHandler)

    fmt.Println("Сервер слушает на порту 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

// Обработчик для /apply
func applyHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Только метод POST поддерживается", http.StatusMethodNotAllowed)
        return
    }

    var s Student
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&s); err != nil {
        http.Error(w, "Невозможно декодировать JSON", http.StatusBadRequest)
        return
    }

    totalScore := s.MathScore + s.InformaticsScore + s.EnglishScore
    if totalScore >= 14 {
        // Добавляем студента в список поступивших
        mu.Lock()
        admittedStudents = append(admittedStudents, s)
        mu.Unlock()
        fmt.Fprintf(w, "Студент %s поступил!\n", s.FullName)
    } else {
        fmt.Fprintf(w, "Студент %s не поступил. Сумма баллов: %d\n", s.FullName, totalScore)
    }
}

// Обработчик для /admitted
func admittedHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    mu.Lock()
    defer mu.Unlock()
    json.NewEncoder(w).Encode(admittedStudents)
}
