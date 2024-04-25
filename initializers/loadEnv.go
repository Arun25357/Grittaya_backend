package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost      string `mapstructure:"POSTGRES_HOST"`
	DBUsername  string `mapstructure:"POSTGRES_USER"`
	DBPassword  string `mapstructure:"POSTGRES_PASSWORD"`
	DBName      string `mapstructure:"POSTGRES_DB"`
	DBPort      string `mapstructure:"POSTGRES_PORT"`
	BackendPort string `mapstructure:"BACKEND_PORT"`

	ClientOrigin string `mapstructure:"FRONTEND_ORIGIN"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
