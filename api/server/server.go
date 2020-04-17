package main

import (
	// This is a named import from another local package. Need for dbconn methods.
	"log"
	"net/http"
	"os"

	gql "api"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	print("connecting to psql")
	ConnectPSQL()
	defer ClosePSQL()

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
