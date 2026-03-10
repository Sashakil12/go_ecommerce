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
	DBConfig    DBConfig
}
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
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
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("dbHost is required")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("dbPort is required")
		os.Exit(1)
	}
	dbPortInt, err := strconv.ParseInt(dbPort, 10, 32)
	if err != nil {
		fmt.Println("dbPort is required")
		os.Exit(1)
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("dbUser is required")
		os.Exit(1)
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("dbPassword is required")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("dbName is required")
		os.Exit(1)
	}
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		fmt.Println("dbSSLMode is required")
		os.Exit(1)
	}
	config.DBConfig = DBConfig{
		Host:     dbHost,
		Port:     int(dbPortInt),
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
		SSLMode:  dbSSLMode,
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
