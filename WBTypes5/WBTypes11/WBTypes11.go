package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type PersonalWb struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

var (
	people []PersonalWb
	log    = logrus.New()
)

func main() {
	//log.Out = os.Stdout
	http.HandleFunc("/homepage", homePageHandler)
	http.HandleFunc("/daysuntill", daysUntilNewYearHandler)
	http.HandleFunc("/personal", peopleHandler)
	http.HandleFunc("/status", statusHandler)

	log.Info("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homePageHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("WB is better that OZON")
	_, err := w.Write([]byte("WB is better that OZON"))
	if err != nil {
		return
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPost:
		postPerson(w, r)
	default:
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL,
		}).Warn("Invalid method")
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to encode people")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Info("Returned list of people")
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	var person PersonalWb
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	people = append(people, person)
	log.WithFields(logrus.Fields{
		"name":     person.Name,
		"position": person.Position,
	}).Info("Added new person")
	w.WriteHeader(http.StatusCreated)
	//w.WriteHeader(http.StatusOK)
}

func statusHandler(w http.ResponseWriter, _ *http.Request) {
	log.Info("Status check")
	w.WriteHeader(http.StatusOK)
	bytesWritten, err := w.Write([]byte("Server is up and running"))
	if err != nil {
		return
	}
	log.Printf("Response written successfully. Bytes written: %d", bytesWritten)
}

func daysUntilNewYearHandler(w http.ResponseWriter, _ *http.Request) {
	currentDate := time.Now()
	// Дата 31 декабря 2024 года
	futureDate := time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC)

	// Разница между датами в днях
	daysUntilNewYear := futureDate.Sub(currentDate).Hours() / 24

	fprintf, err := fmt.Fprintf(w, "Days until January 31, 2024: %.0f\n", daysUntilNewYear)
	if err != nil {
		return
	}
	log.Printf("Response written successfully. Bytes written: %d", fprintf)
}
