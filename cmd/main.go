package main

import (
	"brandsdigger/internal/factory"
	http2 "brandsdigger/internal/interface/http"
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

	// Combine public and protected routes
	router := http2.CreateRouter(authHandler, nameHandler, factory.Token)

	addr := ":8080"
	fmt.Printf("🚀 Starting server on %s\n", addr)
	if err := httpNet.ListenAndServe(addr, router); err != nil {
		log.Fatalf("❌ Error starting server: %v", err)
	}
}
