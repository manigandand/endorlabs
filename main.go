package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	eapi "github.com/manigandand/endorlabs/api"
	"github.com/manigandand/endorlabs/config"
	"github.com/manigandand/endorlabs/store"

	"github.com/go-chi/chi/v5"
	"github.com/manigandand/adk/api"
	appmiddleware "github.com/manigandand/adk/middleware"
	"github.com/rs/cors"
)

var (
	name    = "endor labs"
	version = "1.0.0"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Wrong length of arguments")
	}

	config.Initialize(os.Args[1:]...)

	db := store.Init(config.DBType)
	if db == nil {
		log.Fatalf("🦠store initialize failed 👎")
	}
	api.InitService(name, version)

	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// cross & loger middleware
	router.Use(cors.Handler)
	router.Use(
		appmiddleware.Logger,
		appmiddleware.Recoverer,
	)

	router.Route("/", eapi.Routes)

	interruptChan := make(chan os.Signal, 1)
	go func() {
		signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		// Block until we receive our signal.
		<-interruptChan

		log.Println("Shutting down db...")
		store.Store.Close()
		os.Exit(0)
	}()

	log.Println("Starting server on port:", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router); err != nil {
		log.Fatal(err)
	}
}
