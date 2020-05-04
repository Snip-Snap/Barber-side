package main

import (
	// This is a named import from another local package. Need for dbconn methods.

	"log"
	"net/http"
	"os"

	"api"
	"api/generated"
	"api/internal/auth"
	"api/internal/database"
	"api/internal/directive"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
)

const defaultPort = "69"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	print("Connecting to psql\n")
	database.ConnectPSQL()
	defer database.ClosePSQL()

	c := generated.Config{Resolvers: &api.Resolver{}}
	directive.VerifyAuth(&c)

	server := handler.GraphQL(generated.NewExecutableSchema(c))

	router.Handle("/", handler.Playground("Barbershop", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
