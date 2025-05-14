// Package http provides the server configuration. Including all routes, timeouts and shutdown.
package http

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof for profiling
	"os"
	"os/signal"
	"time"
)

const (
	BASE_PATH = "/kvs"
	PORT      = ":8080"
	// Main server port
)

// StartServer creates an HTTP server that prints path and exposes pprof.
func StartServer(serverStarted chan struct{}, done chan bool) {
	// Handlers
	http.HandleFunc(BASE_PATH+"/ping", Ping)
	http.HandleFunc(BASE_PATH+"/get", Get)
	http.HandleFunc(BASE_PATH+"/add", Add)
	http.HandleFunc(BASE_PATH+"/get_all", GetAll)
	http.HandleFunc(BASE_PATH+"/exists", Exists)
	http.HandleFunc(BASE_PATH+"/count", Count)
	http.HandleFunc(BASE_PATH+"/clear", Clear)
	http.HandleFunc(BASE_PATH+"/delete", Delete)
	http.HandleFunc(BASE_PATH+"/update", Update)
	http.HandleFunc(BASE_PATH+"/upsert", Upsert)

	// Main server
	s := http.Server{
		Addr:         PORT,
		Handler:      http.DefaultServeMux, // Use DefaultServeMux
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		log.Printf("KV Store Listening at %s", PORT)
		close(serverStarted) // Signal that the server has started
		log.Fatal(s.ListenAndServe())
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c // Wait for interrupt signal
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down: %v", err)
		return
	}
	log.Println("Server Stopped")
	done <- true
}
