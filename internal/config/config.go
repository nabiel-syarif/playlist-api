package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	HttpConfig       `mapstructure:"http"`
	Database         `mapstructure:"database"`
	JwtConfig        `mapstructure:"jwt"`
	MailConfig       `mapstructure:"mail_smtp"`
	WorkerPoolConfig `mapstructure:"worker_pool"`
}

type MailConfig struct {
	SmtpHost     string `mapstructure:"smtp_host"`
	SmtpPort     int    `mapstructure:"smtp_port"`
	SmtpUsername string `mapstructure:"smtp_username"`
	SmtpPassword string `mapstructure:"smtp_password"`
}

type WorkerPoolConfig struct {
	WorkerCounts int `mapstructure:"num_of_workers"`
}

type HttpConfig struct {
	Addr            string `mapstructure:"addr"`
	Port            string `mapstructure:"port"`
	GracefulTimeout int    `mapstructure:"graceful_timeout"`
}

type JwtConfig struct {
	SecretKey      string `mapstructure:"secret"`
	ExpirationTime int    `mapstructure:"expires_in"`
}

type Database struct {
	PostgreURL string `mapstructure:"postgre"`
}

func New() (*Config, error) {
	viper.SetConfigType("yaml")
	if os.Getenv("APP_ENV") == "DEVELOPMENT" {
		viper.SetConfigFile("./files/etc/env.development.yaml")
	} else {
		viper.SetConfigFile("./files/etc/env.production.yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
