package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josephspurrier/gocleanarchitecture/cmd/webapp/boot"
)

// main is the entrypoint for the application.
func main() {
	// Set the port.
	port := 8080

	// Register the services and load the routes.
	http.Handle("/", boot.RegisterServices("html").LoadRoutes())

	// Display message on the server.
	log.Printf("webapp started on port %v\n", port)

	// Run the web listener.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
