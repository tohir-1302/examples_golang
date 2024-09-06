package config

import (
	"github.com/spf13/viper"
	"github.com/tohir-1302/gin_jwt_token/helper"
	"time"
)

type Config struct {
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBUsername string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBName     string `mapstructure:"POSTGRES_DB"`

	ServerPort string `mapstructure:"PORT"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAX_AGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	bug := viper.ReadInConfig()
	helper.ErrorPanic(bug)

	err = viper.Unmarshal(&config)
	return
}
