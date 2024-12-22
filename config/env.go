package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Creating a Structure for the configuration details
type Config struct {
	PublicHost string
	Port       string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Environs = initConfig()

// Initializing Database Configuration File Information
func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "gorestapi"),
	}
}

// A function that returns a key-value pair of the config information
func getEnv(key, fallback string) string {
	// Checking to see if the value can be accessed. If it can,
	// it will print out the value
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}
