package main

import (
	"brandsdigger/internal/factory"
	http2 "brandsdigger/internal/interface/http"
	"brandsdigger/internal/interface/http/middleware"
	"brandsdigger/internal/service"
	"fmt"
	"log"
	httpNet "net/http"
)

func main() {
	factory.Init()
	namesService := &service.NamesService{}
	authHandler := http2.NewAuthHandler(namesService)
	nameHandler := http2.NewNamesHandler(namesService)
	publicRouter := http2.PublicRouter(nameHandler)
	authRouter := http2.AuthRouter(authHandler)
	authRouter.Use(middleware.JWTMiddleware(factory.Token))

	// Start the HTTP server
	addr := ":8080"
	fmt.Printf("Starting server on %s\n", addr)
	if err := httpNet.ListenAndServe(addr, nameRouter); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
