package main

import (
	"log"
	"net/http"

	"gocleanarchitecture/lib/boot"
)

// main is the entrypoint for the application.
func main() {
	// Register all services.
	s := boot.RegisterServices()

	// Load all the services into the controller handlers and return the
	// handler with all the routes.
	h := s.LoadRoutes()

	log.Println("Server running...")

	// Run the web listener.
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}
