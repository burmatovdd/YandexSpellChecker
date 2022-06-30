package Viper

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	YANDEXURL string `mapstructure:"YANDEX_URL"`
}

type ConfigViper interface {
	LoadConfig() (config Config, err error)
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("err in config.go: ", err.Error())
	}

	err = viper.Unmarshal(&config)
	return
}
