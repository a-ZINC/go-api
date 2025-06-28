package main

import (
	"api/internal/middleware"
	"api/pkg/utils"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	mux := setupRoutes();
	rl := middleware.NewRateLimiter(10, 1*time.Minute)
	securityMux := rl.RateLimiter(middleware.Compression(middleware.SecurityHeaders(middleware.CorsHeader(mux))))

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	server := &http.Server{
		Addr:         	":8000",
		Handler: 		securityMux,
		TLSConfig:    	tlsConfig,
	}
	go func() {
		log.Printf("Starting server on port %s...", server.Addr)
		key := "key.pem"
		cert := "cert.pem"
		err := server.ListenAndServeTLS(cert, key)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<- quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout((context.Background()), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
		if closeError := server.Close(); closeError != nil {
			log.Printf("Error closing server: %v", closeError)
		}
	} else {
		log.Println("Server gracefully stopped")
	}
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/teachers", teacherHandler)
	mux.HandleFunc("/students", studentHandler)
	mux.HandleFunc("/exams", examHandler)
	return mux
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
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		name := utils.ParseFormData(r);
		log.Println("POST request to /students")
		w.Write([]byte( fmt.Sprintf("Create a new student with name: %s", name)))
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
