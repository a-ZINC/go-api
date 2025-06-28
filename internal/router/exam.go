package router

import (
	"log"
	"net/http"
)

func examHandler(w http.ResponseWriter, r *http.Request) {
	// Handle exam requests
	switch r.Method {
	case http.MethodGet:
		log.Println("GET request to /exams")
		w.Write([]byte("List of exams"))
	case http.MethodPost:
		log.Println("POST request to /exams")
		w.Write([]byte("Create a new exam"))
	case http.MethodPut:
		log.Println("PUT request to /exams")
		w.Write([]byte("Update an exam"))
	case http.MethodDelete:
		log.Println("DELETE request to /exams")
		w.Write([]byte("Delete an exam"))
	case http.MethodPatch:
		log.Println("PATCH request to /exams")
		w.Write([]byte("Partially update an exam"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("Unsupported method: %s", r.Method)
		return
	}
}
