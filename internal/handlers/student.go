package handlers

import (
	"api/pkg/utils"
	"fmt"
	"log"
	"net/http"
)

func studentHandler(w http.ResponseWriter, r *http.Request) {
	// Handle student requests
	switch r.Method {
	case http.MethodGet:
		log.Println("GET request to /students")
		w.Write([]byte("List of students"))
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		name := utils.ParseFormData(r)
		log.Println("POST request to /students")
		w.Write([]byte(fmt.Sprintf("Create a new student with name: %s", name)))
	case http.MethodPut:
		log.Println("PUT request to /students")
		w.Write([]byte("Update a student"))
	case http.MethodDelete:
		log.Println("DELETE request to /students")
		w.Write([]byte("Delete a student"))
	case http.MethodPatch:
		log.Println("PATCH request to /students")
		w.Write([]byte("Partially update a student"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("Unsupported method: %s", r.Method)
		return
	}
}
