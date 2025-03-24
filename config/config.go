package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	fmt.Println("-----------------config-----------------")

	dir, _ := os.Getwd()

	viper.SetConfigType("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(path.Join(dir, "config"))

	// Load environment variables from config.env
	viper.SetConfigName("config.env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Load database configuration if the environment is local
	env := viper.GetString("SERVICE_ENV")
	fmt.Println("ENV: ", env)

	if env == "local" {
		viper.SetConfigName("config.env.local")
		if err := viper.MergeInConfig(); err != nil {
			panic(err)
		}
	}

	// Automatically bind environment variables
	viper.AutomaticEnv()

	fmt.Println("-----------------config-----------------")
}
