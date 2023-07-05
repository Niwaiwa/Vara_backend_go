package configs

import (
	"log"

	"github.com/spf13/viper"
)

// struct to map env values
type Config struct {
	Debug                  bool   `mapstructure:"DEBUG"`
	ServerPort             string `mapstructure:"SERVER_PORT"`
	SecretKey              string `mapstructure:"SECRET_KEY"`
	DBURL                  string `mapstructure:"DB_URL"`
	DBConnMaxLifetimeMs    int32  `mapstructure:"DB_CONN_MAX_LIFETIME_MS"`
	DBMaxOpenConns         int32  `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns         int32  `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBDnsScanIntervalSec   int32  `mapstructure:"DB_DNS_SCAN_INTERVAL_SEC"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRE_HOURS"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRE_HOURS"`
}

// We will call this in main.go to load the env variables and initialize the config variable
func InitConfigs() *Config {
	config := loadEnvVariables()
	return config
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
