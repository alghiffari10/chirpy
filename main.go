package main

import (
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("Server is running on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}
