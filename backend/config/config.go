package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		DSN          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Redis struct {
		Addr     string
		Type     int
		Password string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error reading config file: %v", err)
	}
	AppConfig = &Config{}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("fatal error unmarshal config file: %v", err)
	}
	InitDB()
	InitRedis()
}
