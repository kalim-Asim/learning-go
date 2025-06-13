package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting server on :8080")
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World!</h1>"))
}
