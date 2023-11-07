package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Set-Kaung/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatalln("No port found")
	}

	connString := os.Getenv("DB_URL")
	if connString == "" {
		log.Fatalln("No db connection string")
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	queries := database.New(db)

	apiCnfg := &apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/error", errorHandler)
	v1router.Post("/users", apiCnfg.CreateUser)
	v1router.Get("/users", apiCnfg.GetUser)

	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Starting server on port %s \n", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
