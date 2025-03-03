package main

import (
	"brandsdigger/internal/factory"
	"brandsdigger/internal/infrastructure/http"
	"brandsdigger/internal/service"
	"fmt"
	"log"
	httpNet "net/http"
)

func main() {
	factory.Init()
	nameHandler := http.NewNamesHandler(&service.NamesService{})
	nameRouter := http.NewRouter(nameHandler)

	// Start the HTTP server
	addr := ":8080"
	fmt.Printf("Starting server on %s\n", addr)
	if err := httpNet.ListenAndServe(addr, nameRouter); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
