package main

import (
	"fmt"
	"log"
	"net/http"
)

// The handler function matching the http.HandlerFunc signature
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	if r.Method!=http.MethodGet{
		http.Error(w,"Only GET req is allowed",http.StatusNotAcceptable);
		return;
	}
	_,_=w.Write([]byte(`{"msg":"Hello from Go net/http server"}`));

	// r.URL.Path contains the requested URI path
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	// 1. Create a new custom request multiplexer (router)
	mux := http.NewServeMux()

	// 2. Register patterns and their corresponding handler functions
	mux.HandleFunc("/hello", helloHandler)

	// 3. Start the server on port 8080 using the configured router
	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
