package main

import (
	// This is a named import from another local package. Need for dbconn methods.
	"log"
	"net/http"
	"os"

	"graphqltest/api"
	"graphqltest/api/generated"
	"graphqltest/api/internal/database"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	print("connecting to psql")
	database.ConnectPSQL()
	defer database.ClosePSQL()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &api.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
