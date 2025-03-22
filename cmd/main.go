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
	nameHandler := http2.NewNamesHandler(&service.NamesService{})
	nameRouter := http2.NewRouter(nameHandler)

	// Start the HTTP server
	addr := ":8080"
	fmt.Printf("Starting server on %s\n", addr)
	if err := httpNet.ListenAndServe(addr, nameRouter); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
