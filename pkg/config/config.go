package config

import (
	"fmt"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

func NewConfig() error {
	configPath, err := xdg.ConfigFile("depot/depot.yaml")
	if err != nil {
		return err
	}

	viper.SetConfigFile(configPath)
	viper.SetEnvPrefix("DEPOT")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	return fmt.Errorf("unable to read config file: %v", err)
}

func GetApiToken() string {
	return viper.GetString("api_token")
}

func SetApiToken(token string) error {
	viper.Set("api_token", token)
	return viper.WriteConfig()
}

func ClearApiToken() error {
	viper.Set("api_token", "")
	return viper.WriteConfig()
}
