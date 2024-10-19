package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleware(method string, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w, r)
	}
}

func getMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get", middleware(http.MethodGet, getMessageHandler))

	server := http.Server{
		Addr:    "0.0.0.0:6000",
		Handler: mux,
	}

	// Handling potential error from starting the server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
