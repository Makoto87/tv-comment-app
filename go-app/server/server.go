package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Makoto87/tv-comment-app/go-app/server/graph"
	"github.com/Makoto87/tv-comment-app/go-app/server/graph/generated"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
	}

	origin, okOrigin:= os.LookupEnv("CLIENT_ORIGIN")
	if !okOrigin {
		origin = "http://localhost:3000"
	}

	gqltest := flag.Bool("gqltest", false, "when gqltest flag is used, GraphQL playground is opened")
	flag.Parse()

	if *gqltest {
		fmt.Println("GraphQL playground is used")
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowCredentials: true,
		Debug: true,
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/query", c.Handler(srv))

	httpServer := &http.Server{
		Addr: ":" + serverPort,
	}

	go func() {
		log.Println("GraphQL server is connecting")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("HTTP server Shutdown start")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
}
