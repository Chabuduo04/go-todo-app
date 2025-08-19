package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	JwtSecret []byte
	DBHost    string
	DBPort    int
	DBUser    string
	DBPass    string
	DBName    string
}

var AppConfig *Config

// InitConfig 初始化配置
func InitConfig() {
	// 1. 先尝试加载本地 .env 文件（仅在本地有用，生产容器中无影响）
	_ = godotenv.Load()

	// 2. 从环境变量读取配置
	AppConfig = &Config{
		AppPort:   getEnv("APP_PORT", "8080"), // 普通配置，可有默认值
		JwtSecret: []byte(getEnvRequired("JWT_SECRET")), // 必须显式配置
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnvAsInt("DB_PORT", 3306),
		DBUser:    getEnvRequired("DB_USER"), // 必须显式配置
		DBPass:    getEnvRequired("DB_PASSWORD"), // 必须显式配置
		DBName:    getEnv("DB_NAME", "todo_db"),
	}

}

// getEnv 读取字符串环境变量，支持默认值
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvRequired 读取必须的环境变量（敏感数据），未设置时退出
func getEnvRequired(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("必须设置环境变量 %s", key)
	}
	return value
}


// getEnvAsInt 读取整数环境变量
func getEnvAsInt(key string, defaultVal int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Fatalf("环境变量 %s 必须是整数，但得到: %s", key, valueStr)
		}
		return value
	}
	return defaultVal
}
