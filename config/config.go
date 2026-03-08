package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
}

func loadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("serviceName is required")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	httpPortInt, err := strconv.ParseInt(httpPort, 10, 32)
	if err != nil {
		fmt.Println("httpPort is required")
		os.Exit(1)
	}
	JwtSecret := os.Getenv("JWT_SECRET")
	if JwtSecret == "" {
		fmt.Println("JwtSecret is required")
		os.Exit(1)
	}
	config := &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(httpPortInt),
		JwtSecret:   JwtSecret,
	}
	fmt.Printf("config loaded: %+v\n", config)
	return config
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = loadConfig()
	}
	return config
}
