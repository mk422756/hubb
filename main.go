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
	"github.com/hubbdevelopers/repository"

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

	authClient, err := auth.NewClient()
	if err != nil {
		log.Println("new Authentication Client error")
		log.Println(err.Error())
		panic(err)
	}

	authenticator, err := auth.NewAuthenticator(authClient)
	if err != nil {
		log.Println("new Authentication error")
		log.Println(err.Error())
		panic(err)
	}

	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	resolver := &gql.Resolver{
		UserRepo:      repository.NewUserRepository(db.GetDB()),
		PageRepo:      repository.NewPageRepository(db.GetDB()),
		TagRepo:       repository.NewTagRepository(db.GetDB()),
		Authenticator: authenticator,
	}

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", authenticator.UIDMiddleware(handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: resolver}))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
