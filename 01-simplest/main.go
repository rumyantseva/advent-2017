package main

import (
	"fmt"
	"net/http"
)

// How to try it: go run main.go
func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello! Your request was processed.")
	},
	)
	http.ListenAndServe(":8000", nil)
}
