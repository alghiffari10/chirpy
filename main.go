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

	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Server is running on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}
