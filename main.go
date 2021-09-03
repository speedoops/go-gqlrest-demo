package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/speedoops/go-gqlrest-demo/graph"
	"github.com/speedoops/go-gqlrest-demo/graph/engine"
	"github.com/speedoops/go-gqlrest-demo/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := engine.NewServer(&graph.Resolver{})

	mux := chi.NewRouter()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)
	mux.Handle("/graphql", srv)
	generated.RegisterHandlers(mux, srv, "")

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
