package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)
	apiCfg := &apiConfig{}

	mux.HandleFunc("GET /v1/readiness", apiCfg.handlerReady)
	mux.HandleFunc("GET /v1/err", apiCfg.handlerErr)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Server running on port: %s", port)
	log.Fatal(srv.ListenAndServe())
}
