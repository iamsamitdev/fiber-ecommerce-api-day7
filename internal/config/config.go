package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv       string
	AppPort      string
	AppURL       string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
	JWTExpiresIn string
}

func LoadConfig() (*Config, error) {

	// โหลดไฟล์ .env ถ้ามี
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	config := &Config{

		// ค่าที่ปลอดภัยสำหรับ default
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "3000"),
		AppURL:       getEnv("APP_URL", "http://localhost:3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),

		// ค่าที่ไม่ปลอดภัยสำหรับ default - ต้องกำหนดใน env
		DBPassword: getEnv("DB_PASS", ""),
		DBName:     getEnv("DB_NAME", ""),
		JWTSecret:  getEnv("JWT_SECRET", ""),
	}

	// ตรวจสอบค่าที่จำเป็นต้องมี
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// ฟังก์ชันตรวจสอบค่าที่จำเป็นใน environment variables
func validateConfig(config *Config) error {
	// ตรวจสอบค่าที่จำเป็นสำหรับ production
	if config.AppEnv == "production" {
		if config.DBPassword == "" {
			return fmt.Errorf("DB_PASS is required for production environment")
		}
		if config.JWTSecret == "" {
			return fmt.Errorf("JWT_SECRET is required for production environment")
		}
		if len(config.JWTSecret) < 32 {
			return fmt.Errorf("JWT_SECRET must be at least 32 characters long for production")
		}
		if config.DBSSLMode == "disable" {
			log.Println("Warning: SSL is disabled for database connection in production")
		}
	}

	// ตรวจสอบค่าพื้นฐานที่ต้องมีเสมอ
	if config.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
