package main

import (
	"listenbuddy/internal"
	"log"
	"net/http"
)

func main() {
	err := internal.MakeDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	mux.HandleFunc("/createuser", internal.CreateUser)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
