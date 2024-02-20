package config

import (
	"log"

	"github.com/spf13/viper"
)

type env struct {
	Port     string `mapstructure:"PORT"`
	Host     string `mapstructure:"HOST"`
	CacheDir string `mapstructure:"CACHE_DIR"`
	RembgUrl string `mapstructure:"REMBG_URL"`
}

var Env *env

func InitEnv() {
	Env = loadConfig()
}

func loadConfig() (e *env) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&e)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return
}
