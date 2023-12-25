package config

import (
	"fmt"
	"log"
	"os"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/logger"
	"github.com/spf13/viper"
)

type AppConfig struct {
	PORT int16
}

var AppConfigInstance AppConfig

func LoadConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	envFile := fmt.Sprintf("%s/.env", dir)
	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&AppConfigInstance); err != nil {
		logger.Fatal(err.Error())
	}
}

func BootstrapApp() {
	logger.Init()
	LoadConfig()
}
