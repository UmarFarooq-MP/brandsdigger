package main

import (
	"brandsdigger/internal/factory"
	lhttp "brandsdigger/internal/interface/http"
	"brandsdigger/internal/service"
	"fmt"
	"log"
	httpNet "net/http"
)

func main() {
	factory.Init()

	namesService := &service.NamesService{}
	authHandler := lhttp.NewAuthHandler(namesService)
	nameHandler := lhttp.NewNamesHandler(namesService)

	// Combine public and protected routes
	router := lhttp.CreateRouter(authHandler, nameHandler, factory.Token)

	addr := ":8080"
	fmt.Printf("Starting server on %s\n", addr)
	if err := httpNet.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
