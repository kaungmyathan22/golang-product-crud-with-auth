package config

import (
	"fmt"
	"log"
	"os"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/logger"
	"github.com/spf13/viper"
)

type AppConfig struct {
	PORT               string
	DATABASE_URI       string
	DATABASE_NAME      string
	USER_COLLECTION    string
	PRODUCT_COLLECTION string
}

var AppConfigInstance AppConfig

func LoadConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envFile := fmt.Sprintf("%s/.env", dir)
	viper.SetConfigFile(envFile)
	viper.ReadInConfig()
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&AppConfigInstance); err != nil {
		logger.Fatal(err.Error())
	}
}

func BootstrapApp() {
	logger.Init()
	LoadConfig()
}
