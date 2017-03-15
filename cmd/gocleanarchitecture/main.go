package main

import (
	"log"
	"net/http"

	"github.com/josephspurrier/gocleanarchitecture/lib/boot"
)

// main is the entrypoint for the application.
func main() {
	// Register the services and load the routes.
	http.Handle("/", boot.RegisterServices().LoadRoutes())

	// Display message on the server.
	log.Println("Server started.")

	// Run the web listener.
	http.ListenAndServe(":8080", nil)
}
