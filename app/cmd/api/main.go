package main

import (
	"context"
	"github/thomasfurland/go-simple-webservice/internal/handlers"
	"github/thomasfurland/go-simple-webservice/internal/httpserver"
	"log"
	"os"
	"os/signal"
	"time"
)

const Port = ":8080"

func main() {
	mux := handlers.New()
	server := httpserver.New(Port, mux, httpserver.Options{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := httpserver.Run(ctx, server); err != nil {
		log.Fatal(err)
	}
}
