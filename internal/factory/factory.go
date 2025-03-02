package factory

import (
	"brandsdigger/internal/client"
	"brandsdigger/internal/client/openai"
	"brandsdigger/internal/config"
	"github.com/spf13/viper"
)

// Global variable, but uninitialized initially
var Generate client.MessagesGenerator

// Init function to initialize Generate
func Init() {
	config.Init()
	Generate = openai.New(
		viper.GetString(config.OPENAI_API_KEY),
		viper.GetString(config.OPENAI_MODEL_NAME),
	)
}
