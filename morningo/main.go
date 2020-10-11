package main

import (
	"fmt"
	"log"
	"net/http"
)

// handleHello GET /hello
func handleHello(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method, r.RequestURI)

	// Returns hello world! as a response
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	// registers handleHello to GET /hello
	http.HandleFunc("/api/hello", handleHello)
	// starts the server on port 5000
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalln(err)
	}
}