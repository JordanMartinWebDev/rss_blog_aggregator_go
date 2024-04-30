package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jordanmatinwebdev/rss_blog_agregator_go/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	mux := http.NewServeMux()
	corsMux := middlewareCors(mux)

	apiCfg := &apiConfig{
		DB: dbQueries,
	}

	//Users
	mux.HandleFunc("POST /v1/users", apiCfg.handlerCreateUsers)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))

	//Feeds
	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeeds))
	mux.HandleFunc("GET /v1/feeds", apiCfg.handlerGetAllFeeds)

	//Ready
	mux.HandleFunc("GET /v1/readiness", apiCfg.handlerReady)
	mux.HandleFunc("GET /v1/err", apiCfg.handlerErr)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Server running on port: %s", port)
	log.Fatal(srv.ListenAndServe())
}
