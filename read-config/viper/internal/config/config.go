package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config Config

func GetConfig() Config {
	return config
}

type Config struct {
	DB Database `mapstructure:"database"`
}
type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func InitConfig() {
	// 初始化 Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")

	// 设置默认值
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.username", "root")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 将配置解析到结构体
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	// 使用配置
	fmt.Printf("Database config: host: %s, port: %d, username: %s, password: %s\n",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password)

	// 监听配置变更
	viper.WatchConfig()
}
