package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	GODADDY_API_KEY    = "godaddy.api.key"
	GODADDY_API_SECRET = "godaddy.api.secret"
	GODADDY_BASE_URL   = "godaddy.base.url"

	DEEPSEEK_API_KEY    = "deepseek.api.key"
	DEEPSEEK_BASE_URL   = "deepseek.api.url"
	DEEPSEEK_MODEL_NAME = "deepseek.model.name"

	OPENAI_API_KEY    = "openai.api.key"
	OPENAI_MODEL_NAME = "openai.model.name"
)

func Init() {
	fmt.Println("RAW OS ENV:", os.Getenv("API_KEY"), os.Getenv("API_SECRET"), os.Getenv("BASE_URL"))
	_ = viper.BindEnv(GODADDY_API_KEY, "GODADDY_API_KEY")
	_ = viper.BindEnv(GODADDY_API_SECRET, "GODADDY_API_SECRET")
	_ = viper.BindEnv(GODADDY_BASE_URL, "GODADDY_BASE_URL")

	_ = viper.BindEnv(DEEPSEEK_API_KEY, "DEEPSEEK_API_KEY")
	_ = viper.BindEnv(DEEPSEEK_BASE_URL, "DEEPSEEK_BASE_URL")
	_ = viper.BindEnv(DEEPSEEK_MODEL_NAME, "DEEPSEEK_MODEL_NAME")

	_ = viper.BindEnv(OPENAI_API_KEY, "OPENAI_API_KEY")

	viper.SetDefault(DEEPSEEK_API_KEY, "sk-85d4b7c15a1845852aaf947d")
	viper.SetDefault(DEEPSEEK_MODEL_NAME, "deepseek-chat")
	viper.SetDefault(DEEPSEEK_BASE_URL, "https://api.deepseek.com/v1/chat/completions")

	viper.SetDefault(OPENAI_MODEL_NAME, "gpt-3.5-turbo")
	viper.SetDefault(GODADDY_BASE_URL, "https://api.ote-godaddy.com")

}
