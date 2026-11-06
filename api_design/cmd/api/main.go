package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK)
	
	w.Write([]byte(`{"Status":"Ok","Message":"Api working good"}`))
}

func main() {
	mux:=http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)

	err := http.ListenAndServe(":7070", mux)
	if err != nil {
		log.Fatalf("Server Failed:: %v", err)
	}
}
