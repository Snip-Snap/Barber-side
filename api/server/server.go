package main

import (
	// This is a named import from another local package. Need for dbconn methods.
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"api"
	"api/generated"
	"api/internal/auth"
	"api/internal/database"

	"github.com/99designs/gqlgen/graphql"
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
	c.Directives.CheckAuth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		barber := auth.ForContext(ctx)

		if barber != nil {
			if barber.UserName != "" {
				// Let barber proceed with api calls.
				return next(ctx)
			}
		}
		return nil, errors.New("Unauthorised")
	}

	server := handler.GraphQL(generated.NewExecutableSchema(c))

	router.Handle("/", handler.Playground("Barbershop", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
