package config

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	dir, _ := os.Getwd()

	// Automatically bind environment variables
	viper.AutomaticEnv()

	// Set the configuration type and path
	viper.SetConfigType("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(path.Join(dir, "config"))
	viper.ReadInConfig()

	env := viper.GetString("SERVICE_ENV")
	if env == "" {
		env = "local"
	}
	log.Println("SERVICE_ENV: ", env)

	// Load local environment variables via config.env.local
	if env == "local" {
		viper.SetConfigName("config.env.local")
		if err := viper.MergeInConfig(); err != nil {
			panic(err)
		}
	}

}
