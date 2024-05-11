package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBPassword          string        `mapstructure:"DB_PASSWORD"`
	DBPort              string        `mapstructure:"DB_PORT"`
	DBHost              string        `mapstructure:"DB_HOST"`
	DBName              string        `mapstructure:"DB_NAME"`
	DBUser              string        `mapstructure:"DB_USER"`
	DBSSLMode           string        `mapstructure:"DB_SSL_MODE"`
	TokenSyemmetricKey  string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	Port                string        `mapstructure:"PORT"`
}

// load config from .env file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
