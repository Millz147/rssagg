package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/millz147/rssagg/internal/database"
)

type dbRepoStructure struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT can not be found in the environment.")
	}

	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		log.Fatal("DB URI can not be found in the environment.")
	}

	conn, err := sql.Open("postgres", dbUri)
	if err != nil {
		log.Fatal("Database is not connected.")
	}

	queries := database.New(conn)
	dbRepo:= dbRepoStructure{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	routerV1 := chi.NewRouter()

	routerV1.Get("/healthz", handlerRediness)
	routerV1.Get("/err", handlerError)
	routerV1.Post("/create-user",dbRepo.handleCreateUser)
	router.Mount("/v1", routerV1)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Sever listening on port %v", port)
	err =srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PORT", port)

}
