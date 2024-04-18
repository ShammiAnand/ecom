package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost  string
	Port        string
	DBUser      string
	DBPassword  string
	DBAddress   string
	DBName      string
	Secret      string
	TokenExpiry int
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()

	tokenExpiry := getEnv("TOKEN_EXPIRY", 3600).(int)

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost").(string),
		Port:       getEnv("PORT", "8000").(string),
		DBUser:     getEnv("DB_USER", "root").(string),
		DBPassword: getEnv("DB_PASSWORD", "mypassword").(string),
		DBName:     getEnv("DB_NAME", "ecom").(string),
		DBAddress: fmt.Sprintf(
			"%s:%s",
			getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306"),
		),
		Secret:      getEnv("SECRET", "some_secret").(string),
		TokenExpiry: tokenExpiry,
	}
}

func getEnv(key string, fallback interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		switch fallback.(type) {
		case string:
			return value
		case int:
			if intValue, err := strconv.Atoi(value); err == nil {
				return intValue
			}
		case float64:
			if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
				return floatValue
			}
		default:
			return value
		}
	}
	return fallback
}
