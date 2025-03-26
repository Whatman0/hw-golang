package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email    EmailConfig
	Password PasswordConfig
	Addr     AddressConfig
}

type EmailConfig struct {
	Email string
}

type PasswordConfig struct {
	Password string
}

type AddressConfig struct {
	Addr string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Email: EmailConfig{
			Email: os.Getenv("EMAIL"),
		},
		Password: PasswordConfig{
			Password: os.Getenv("PASSWORD"),
		},
		Addr: AddressConfig{
			Addr: os.Getenv("ADDR"),
		},
	}
}
