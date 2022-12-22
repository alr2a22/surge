package config

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisPort     string `mapstructure:"REDIS_PORT"`

	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"POSTGRES_DB"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`

	LogLevel string `mapstructure:"LOGLEVEL"`

	JwtSecret string `mapstructure:"JWT_SECRET"`
}

var cfg *Config
var onceConfig sync.Once

func load() {
	viper.MustBindEnv("REDIS_HOST")
	viper.MustBindEnv("REDIS_PASSWORD")
	viper.MustBindEnv("REDIS_PORT")
	viper.MustBindEnv("POSTGRES_USER")
	viper.MustBindEnv("POSTGRES_PASSWORD")
	viper.MustBindEnv("POSTGRES_DB")
	viper.MustBindEnv("POSTGRES_HOST")
	viper.MustBindEnv("POSTGRES_PORT")
	viper.MustBindEnv("LOGLEVEL")
	viper.MustBindEnv("JWT_SECRET")

	err := viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infoln("configuration loaded successfully:", fmt.Sprintf("%+v\n", cfg))
}

func GetConfig() *Config {
	onceConfig.Do(func() {
		load()
	})
	return cfg
}
