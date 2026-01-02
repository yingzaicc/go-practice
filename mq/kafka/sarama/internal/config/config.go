package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var config Config

func GetConfig() Config {
	return config
}

type Config struct {
	Brokers  []string `yaml:"brokers"`
	Version  string   `yaml:"version"`
	Producer Producer `yaml:"producer"`
	Consumer Consumer `yaml:"consumer"`
}

type Producer struct {
	Count int    `yaml:"count"`
	Topic string `yaml:"topic"`
}

type Consumer struct {
	Topic []string `yaml:"topic"`
}

func InitConfig() {
	file := filepath.Join("../../config", "config.yaml")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic("config file not found")
	}

	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	fmt.Printf("Kafka config: brokers: %v, producer: %v, consumer: %v\n",
		config.Brokers, config.Producer, config.Consumer)
}
