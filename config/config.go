// Package config loads environment variables
package config

import "github.com/spf13/viper"

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	AppName       string `mapstructure:"APP_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		config = Config{
			ServerAddress: "0.0.0.0:8080",
			AppName:       "APP",
		}
		return config, nil
	}
	err = viper.Unmarshal(&config)
	return
}
