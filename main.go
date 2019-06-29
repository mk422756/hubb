package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/hubbdevelopers/auth"
	"github.com/hubbdevelopers/db"
	"github.com/hubbdevelopers/gql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	db.Connect()
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", auth.UIDMiddleware(handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}}))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
