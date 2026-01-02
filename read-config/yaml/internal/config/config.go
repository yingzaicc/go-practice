package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var config Config

func GetConfig() Config {
	return config
}

type Config struct {
	DB Database `yaml:"database"`
}
type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func InitConfig() {
	// 构建配置文件路径
	configPath := filepath.Join("../../config", "config.yaml")

	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Printf("Config file not found: %s\n", configPath)
		return
	}

	// 读取配置文件
	file, err := os.OpenFile(configPath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// 将配置解析到结构体
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	// 使用配置
	fmt.Printf("Database config: host: %s, port: %d, username: %s, password: %s\n",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password)

}
