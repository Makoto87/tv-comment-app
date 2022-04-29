package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	httpServer := &http.Server{
		Addr: ":" + port,
	}

	go func() {
		fmt.Println("GraphQL server is connecting")
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	fmt.Println("HTTP server Shutdown start")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(timeoutCtx); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
}
