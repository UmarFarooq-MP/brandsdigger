package factory

import (
	"brandsdigger/internal/config"
	"brandsdigger/internal/domain/auth"
	"brandsdigger/internal/domain/dns"
	"brandsdigger/internal/domain/names"
	"brandsdigger/internal/infrastructure/client/godaddy"
	"brandsdigger/internal/infrastructure/client/openai"
	lJwt "brandsdigger/internal/infrastructure/jwt"
	"github.com/spf13/viper"
)

// Generate Global variable, but uninitialized initially
var (
	Generate        names.Generator
	DomainValidator dns.DomainValidator
	Token           *auth.TokenService
)

// Init function to initialize Generate
func Init() {

	config.Init()
	Generate = openai.New(
		viper.GetString(config.OPENAI_API_KEY),
		viper.GetString(config.OPENAI_MODEL_NAME),
	)
	DomainValidator = godaddy.New(
		viper.GetString(config.GODADDY_API_KEY),
		viper.GetString(config.GODADDY_API_SECRET),
		viper.GetString(config.GODADDY_BASE_URL),
	)
	Token = lJwt.New("Shahroz_is_gay")
}
