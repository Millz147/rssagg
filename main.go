package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT can not be found in the environment.")
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
  	routerV1:= chi.NewRouter()

	routerV1.Get("/healthz",handlerRediness)
	routerV1.Get("/err",handlerError)
	router.Mount("/v1",routerV1)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Sever listening on port %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PORT", port)

}
