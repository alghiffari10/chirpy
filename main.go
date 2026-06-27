package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var port = ":8080"

type apiConfig struct {
	fileServerHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) metricHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	hits := cfg.fileServerHits.Load()
	message := fmt.Sprintf("Hits: %v", hits)

	w.Write([]byte(message))

}

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	var reset int32 = 0
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	cfg.fileServerHits.Store(reset)

	w.Write([]byte("Reset the metrics"))
}

func main() {
	apiCfg := apiConfig{}
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	mux.Handle("/app/",
		http.StripPrefix("/app/",
			apiCfg.middlewareMetricsInc(http.FileServer(http.Dir("."))),
		),
	)

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/metrics", apiCfg.metricHandler)

	mux.HandleFunc("/reset", apiCfg.resetHandler)

	log.Printf("Server is running on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}
