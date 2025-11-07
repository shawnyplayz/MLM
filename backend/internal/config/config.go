package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	CORS     CORSConfig
	MLM      MLMConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

type CORSConfig struct {
	Origins []string
}

type MLMConfig struct {
	DefaultTreeType           string
	BinaryMaxWidth            int
	MatrixWidth               int
	MatrixDepth               int
	DirectReferralCommission  float64
	LevelCommissionPercentage float64
	MaxCommissionLevels       int
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	jwtExpiry, err := time.ParseDuration(getEnv("JWT_EXPIRY", "24h"))
	if err != nil {
		jwtExpiry = 24 * time.Hour
	}

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "root"),
			DBName:   getEnv("DB_NAME", "mlm_app"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-secret-key"),
			Expiry: jwtExpiry,
		},
		CORS: CORSConfig{
			Origins: strings.Split(getEnv("CORS_ORIGINS", "http://localhost:3000"), ","),
		},
		MLM: MLMConfig{
			DefaultTreeType:           getEnv("DEFAULT_TREE_TYPE", "binary"),
			BinaryMaxWidth:            getEnvAsInt("BINARY_MAX_WIDTH", 2),
			MatrixWidth:               getEnvAsInt("MATRIX_WIDTH", 3),
			MatrixDepth:               getEnvAsInt("MATRIX_DEPTH", 9),
			DirectReferralCommission:  getEnvAsFloat("DIRECT_REFERRAL_COMMISSION", 10.0),
			LevelCommissionPercentage: getEnvAsFloat("LEVEL_COMMISSION_PERCENTAGE", 5.0),
			MaxCommissionLevels:       getEnvAsInt("MAX_COMMISSION_LEVELS", 10),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsFloat(key string, defaultValue float64) float64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultValue
}
