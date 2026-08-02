package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

func home(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://localhost:8081/employees")
	if err != nil {
		http.Error(w, "Employee Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Unable to read response", http.StatusInternalServerError)
		return
	}

	var employees []Employee

	err = json.Unmarshal(body, &employees)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, employees)
	if err != nil {
		http.Error(w, "Template Execute Error", http.StatusInternalServerError)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gateway OK"))
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)

	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}