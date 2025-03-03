package main

import (
	"brandsdigger/internal/factory"
	"fmt"
)

func main() {
	factory.Init()
	messages, err := factory.Generate.GenerateNames("I am a potato Company")
	if err != nil {
		panic(err)
	}
	res, err := factory.DomainValidator.ValidateDomain(messages)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
