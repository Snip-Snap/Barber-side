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

	print("connecting to psql")
	database.ConnectPSQL()
	defer database.ClosePSQL()

	server := handler.GraphQL(generated.NewExecutableSchema(
		generated.Config{Resolvers: &api.Resolver{}}))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
