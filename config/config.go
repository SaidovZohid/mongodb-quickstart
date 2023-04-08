package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	MongoDB MongoDB
}

type MongoDB struct {
	Url      string
	Username string
	Password string
	Database string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env")

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		MongoDB: MongoDB{
			Url:      conf.GetString("MONGODB_URL"),
			Username: conf.GetString("MONGODB_USERNAME"),
			Password: conf.GetString("MONGODB_PASSWORD"),
			Database: conf.GetString("MONGODB_DATABASE"),
		},
	}

	return cfg
}
