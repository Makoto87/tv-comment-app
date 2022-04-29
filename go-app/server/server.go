package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Makoto87/tv-comment-app/go-app/server/graph"
	"github.com/Makoto87/tv-comment-app/go-app/server/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gqltest := flag.Bool("gqltest", false, "when gqltest flag is used, GraphQL playground is opened")
	flag.Parse()

	if *gqltest {
		fmt.Println("GraphQL playground is used")
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/query", srv)

	fmt.Println("GraphQL server is connected")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
