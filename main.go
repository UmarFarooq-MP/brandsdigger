package main

import (
	"brandsdigger/internal/factory"
	"fmt"
)

func main() {
	factory.Init()
	fmt.Println(factory.Generate.GenerateNames("I am a potato Company"))
}
