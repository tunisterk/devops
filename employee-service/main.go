package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var employees = []Employee{
	{1, "Alice", "DevOps"},
	{2, "Bob", "Backend"},
	{3, "Charlie", "QA"},
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Employee Service OK"))
}

func main() {

	http.HandleFunc("/employees", employeeHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Employee Service running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}