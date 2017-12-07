package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rumyantseva/advent-2017/07-makefile/handlers"
)

// How to try it: PORT=8000 go run main.go
func main() {
	log.Print("Starting the service...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
