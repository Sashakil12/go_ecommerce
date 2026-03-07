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
}

func loadConfig() Config {
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
	cnf := Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(httpPortInt),
	}
	fmt.Printf("config loaded: %+v\n", cnf)
	return cnf
}

var config Config = loadConfig()

func GetConfig() *Config {
	return &config
}
