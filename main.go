package main

import (
	"listenbud/middleware"
	"log"
	"net/http"
)

func main() {
	err := middleware.MakeDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	mux.HandleFunc("/createuser", middleware.CreateUser)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
