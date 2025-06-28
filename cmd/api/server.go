package main

import (
	"api/internal/middleware"
	"api/internal/router"
	"context"
	"crypto/tls"
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

	mux := router.SetupRoutes()
	rl := middleware.NewRateLimiter(10, 1*time.Minute)
	hpp := middleware.NewHPP(true, true, "application/x-www-form-urlencoded", []string{"name", "age", "email"})
	securityMux := hpp.HPP()(rl.RateLimiter(middleware.Compression(middleware.SecurityHeaders(middleware.CorsHeader(mux)))))

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	server := &http.Server{
		Addr:      ":8000",
		Handler:   securityMux,
		TLSConfig: tlsConfig,
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

	<-quit
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
