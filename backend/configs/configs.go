package configs

import (
	"os"

	"github.com/joho/godotenv"
)

// DB 관련 설정값 구조체
type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
	TimeZone string
}

// API 관련 설정값 구조체
type Config struct {
	DB DBConfig
	ServerAddr string
	ApiMode string
	ProductionDomain string
	TestDomain string
	JWTSecret []byte
}

// 사전 설정 호출
func Load() *Config {
	_ = godotenv.Load() // 로컬 개발용
	
	cfg := &Config{
		ServerAddr: os.Getenv("API_ADDR") + ":" + os.Getenv("API_PORT"),
		ApiMode: os.Getenv("API_MODE"),
		ProductionDomain: os.Getenv("API_PRODUCTION_DOMAIN_CORS"),
		TestDomain: os.Getenv("API_TEST_DOMAIN_CORS"),
		JWTSecret: []byte(os.Getenv("API_JWT_SECRET")),
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

