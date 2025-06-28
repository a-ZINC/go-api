package handlers

import (
	"log"
	"net/http"
)

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("GET request to /teachers")
		w.Write([]byte("List of teachers"))
	case http.MethodPost:
		log.Println("POST request to /teachers")
		w.Write([]byte("Create a new teacher"))
	case http.MethodPut:
		log.Println("PUT request to /teachers")
		w.Write([]byte("Update a teacher"))
	case http.MethodDelete:
		log.Println("DELETE request to /teachers")
		w.Write([]byte("Delete a teacher"))
	case http.MethodPatch:
		log.Println("PATCH request to /teachers")
		w.Write([]byte("Partially update a teacher"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("Unsupported method: %s", r.Method)
		return
	}
}
