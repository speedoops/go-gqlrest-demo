package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/speedoops/gql2rest/handlerx"
	"github.com/vektah/gqlgen-todos/graph"
	"github.com/vektah/gqlgen-todos/graph/generated"
	"github.com/vektah/gqlgen-todos/graph/model"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		if !getCurrentUser(ctx).HasRole(role) {
			// block calling the next resolver
			return nil, fmt.Errorf("Access denied")
		}
		log.Println("hasRole")

		// or let it pass through
		return next(ctx)
	}

	c.Directives.Hide = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		forArg []string) (res interface{}, err error) {
		return next(ctx)
	}
	c.Directives.Http = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		url string, method *string) (res interface{}, err error) {
		return next(ctx)
	}
	c.Directives.Preview = func(ctx context.Context, obj interface{}, next graphql.Resolver,
		toggledBy string) (res interface{}, err error) {
		return next(ctx)
	}

	srv := handlerx.NewDefaultServer(generated.NewExecutableSchema(c))

	r := chi.NewRouter()
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)
	// r.Handle("/graphql", srv)

	generated.RegisterHandlers("/api/v1", r, srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getCurrentUser(ctx context.Context) model.User {
	return model.User{}
}
