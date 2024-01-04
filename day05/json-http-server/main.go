package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"main.go/internal/server"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlePost).Methods("POST")
	r.HandleFunc("/", handleGet).Methods("GET")

	srv := server.NewHTTPServer(":8080")
	srv.ListenAndServe()
}

// I’ll need to write handlers for them
func handleGet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get\n")
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "post\n")
}
