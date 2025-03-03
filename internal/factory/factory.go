package factory

import (
	"brandsdigger/internal/client"
	"brandsdigger/internal/client/daddy"
	"brandsdigger/internal/client/openai"
	"brandsdigger/internal/config"
	"github.com/spf13/viper"
)

// Generate Global variable, but uninitialized initially
var Generate client.MessagesGenerator
var DomainValidator client.DomainValidator

// Init function to initialize Generate
func Init() {

	config.Init()
	Generate = openai.New(
		viper.GetString(config.OPENAI_API_KEY),
		viper.GetString(config.OPENAI_MODEL_NAME),
	)
	DomainValidator = daddy.New(
		viper.GetString(config.GODADDY_API_KEY),
		viper.GetString(config.GODADDY_API_SECRET),
		viper.GetString(config.GODADDY_BASE_URL),
	)
}
