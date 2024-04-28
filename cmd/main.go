package main

import (
	"log"
	"net/http"

	"github.com/jsatof/listenbuddy/internal"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	mux.HandleFunc("/createuser", internal.CreateUser)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
