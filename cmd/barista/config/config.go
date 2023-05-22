package config

import (
	"fmt"
	configs "go-coffeeshop/pkg/config"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		configs.App  `yaml:"app"`
		configs.HTTP `yaml:"http"`
		configs.Log  `yaml:"logger"`
		RabbitMQ     `yaml:"rabbit_mq"`
	}

	RabbitMQ struct {
		URL string `env-required:"true" yaml:"url" env:"RABBITMQ_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// debug
	fmt.Println(dir)

	err = cleanenv.ReadConfig(dir+"/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
