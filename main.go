package main

import (
	"fmt"
	"log"
	"net/http"
)

type Output struct {
	Message string `json:"message"`
}

// return Hello World JSON
func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello, World!!"}`)
	return
}

// This function is main routine.
func main() {
	fmt.Println("Hello, World!!!")

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
