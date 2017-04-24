package main

import (
	"log"
	"net/http"
)

// main is the entrypoint for the application.
func main() {
	// Register the services and load the routes.
	http.Handle("/", RegisterServices("html").LoadRoutes())

	// Display message on the server.
	log.Println("Server started.")

	// Run the web listener.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
