package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
	TimeZone string
}

type Config struct {
	DB DBConfig
	ServerAddr string
}

// 사전 설정 호출
func Load() *Config {
	_ = godotenv.Load() // 로컬 개발용
	
	cfg := &Config{
		ServerAddr: os.Getenv("API_ADDR") + ":" + os.Getenv("API_PORT"),
		DB: DBConfig{
			Host : os.Getenv("DB_HOST"),
			Port : os.Getenv("DB_PORT"),
			User : os.Getenv("DB_USER"),
			Password : os.Getenv("DB_PASSWORD"),
			DBName : os.Getenv("DB_DATABASE"),
			SSLMode : os.Getenv("DB_SSLMODE"),
			TimeZone : os.Getenv("DB_TIMEZONE"),
		},
	}

	return cfg
}

