package configs

import (
	"log"

	"github.com/spf13/viper"
)

// We will call this in main.go to load the env variables and initialize the config variable
func InitConfigs() *Config {
	config := loadEnvVariables()
	return config
}

// struct to map env values
type Config struct {
	Debug      bool   `mapstructure:"DEBUG"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	SecretKey  string `mapstructure:"SECRET_KEY"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

// Call to load the variables from env
func loadEnvVariables() (config *Config) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(".")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return config
}
