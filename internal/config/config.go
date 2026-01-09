package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Env      string        `mapstructure:"Env"`
	Server   Server        `mapstructure:"Server"`
	Broker   MessageBroker `mapstructure:"Kafka"`
	Redis    Redis         `mapstructure:"Redis"`
	Database Database      `mapstructure:"Postgres"`
}

type Server struct {
	Addr     string        `mapstructure:"Addr"`
	TimeOut  time.Duration `mapstructure:"timeout"`
	IdleTime time.Duration `mapstructure:"idletimeout"`
}

type MessageBroker struct {
	ConnString string `mapstructure:"connString"`
}

type Redis struct {
}

type Database struct {
	ConnString string `mapstructure:"connString"`
}

func MustLoad() *Config {

	cfg, err := getConfigPathFromFlag()
	if err != nil {
		log.Printf("Error read config from patch%s\n", err)

		cfg, err = getConfigFromEnv()

		if err != nil {
			log.Panicf("Error load config from env: %s\n", err)
		}
	}

	return cfg
}

func getConfigPathFromFlag() (*Config, error) {

	var cfg Config

	pflag.String("config", "", "Set patch to config file")

	pflag.Parse()

	cfgStr, _ := pflag.CommandLine.GetString("config")

	viper.SetConfigFile(cfgStr)

	var cfgErr viper.ConfigFileNotFoundError

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &cfgErr) {
			return nil, fmt.Errorf("Error config file not found%s", err)
		}
		return nil, fmt.Errorf("Error read in config %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicf("Error unmarshal config:%s", err)
	}

	return &cfg, nil

}

func getConfigFromEnv() (*Config, error) {

	var cfg Config

	os.Setenv("CONFIG_PATH", "../NotesServices/config/config.yaml")

	viper.SetEnvPrefix("config")
	viper.BindEnv("path")

	path := viper.GetString("path")

	viper.SetConfigFile(path)

	var cfgErr viper.ConfigFileNotFoundError

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &cfgErr) {
			return nil, fmt.Errorf("Error config file not found%s", err)
		}
		return nil, fmt.Errorf("Error read in config %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Error unmarshal config:%s", err)
	}

	return &cfg, nil
}
