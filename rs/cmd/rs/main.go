package main

import (
	"fmt"
	"log"
	"net/http"
	"rs/config"
	"rs/internal/graph/generated"
	"rs/internal/graph/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// init config
	cfg, err := config.Init()
	if err != nil {
		fmt.Errorf("fail to init config: %w\n", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	if cfg.IsLocal() {
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
		http.Handle("/graphql", srv)
	}

	port := cfg.Port
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
