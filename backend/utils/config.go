package utils

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBPassword           string        `mapstructure:"DB_PASSWORD"`
	DBPort               string        `mapstructure:"DB_PORT"`
	DBHost               string        `mapstructure:"DB_HOST"`
	DBName               string        `mapstructure:"DB_NAME"`
	DBUser               string        `mapstructure:"DB_USER"`
	DBSSLMode            string        `mapstructure:"DB_SSL_MODE"`
	TokenSyemmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	Port                 string        `mapstructure:"PORT"`
	ADMIN_EMAIL          string        `mapstructure:"ADMIN_EMAIL"`
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
	os.Setenv("PORT", config.Port)
	return
}

// load config from environment variables
func LoadConfigFromEnv() (config Config) {
	config = Config{
		DBDriver:             viper.GetString("DB_DRIVER"),
		DBPassword:           viper.GetString("DB_PASSWORD"),
		DBPort:               viper.GetString("DB_PORT"),
		DBHost:               viper.GetString("DB_HOST"),
		DBName:               viper.GetString("DB_NAME"),
		DBUser:               viper.GetString("DB_USER"),
		DBSSLMode:            viper.GetString("DB_SSL_MODE"),
		TokenSyemmetricKey:   viper.GetString("TOKEN_SYMMETRIC_KEY"),
		AccessTokenDuration:  viper.GetDuration("ACCESS_TOKEN_DURATION"),
		RefreshTokenDuration: viper.GetDuration("REFRESH_TOKEN_DURATION"),
		Port:                 viper.GetString("PORT"),
		ADMIN_EMAIL:          viper.GetString("ADMIN_EMAIL"),
	}
	return
}
