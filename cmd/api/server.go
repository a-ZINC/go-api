package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	serverAddr := ":8080"
	http.HandleFunc("/teachers", teacherHandler)
	http.HandleFunc("/students", studentHandler)
	http.HandleFunc("/exams", examHandler)
	isActive := make(chan os.Signal, 1)
	signal.Notify(isActive, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		file, err := os.OpenFile("./shutdown.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		ticker := time.NewTicker(5 * time.Second)
		if err != nil {
			log.Fatalf("Error creating shutdown log file: %v", err)
		}
		defer file.Close()
		for {
			select {
			case <-isActive:
				log.Println("Received shutdown signal, shutting down server...")
				_, err = file.WriteString("Server shutdown initiated\n")
				if err != nil {
					log.Fatalf("Error writing to shutdown log file: %v", err)
				}
				os.Exit(0)
			case <-ticker.C:
				log.Println("Server is running on", serverAddr)

			}
		}

	}()
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

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

func studentHandler(w http.ResponseWriter, r *http.Request) {
	// Handle student requests
	switch r.Method {
	case http.MethodGet:
		log.Println("GET request to /students")
		w.Write([]byte("List of students"))
	case http.MethodPost:
		log.Println("POST request to /students")
		w.Write([]byte("Create a new student"))
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
