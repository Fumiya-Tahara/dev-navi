package postgresql

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Username string
	Password string
	DbName   string
	Port     string
	SslMode  string
	TimeZone string
}

func NewConfig() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	host := viper.GetString("POSTGRES_HOST")
	userName := viper.GetString("POSTGRES_USER")
	password := viper.GetString("POSTGRES_PASSWORD")
	dbName := viper.GetString("POSTGRES_DB")
	port := viper.GetString("POSTGRES_PORT")
	sslMode := viper.GetString("POSTGRES_SSLMODE")
	timeZone := viper.GetString("POSTGRES_TIMEZONE")

	cfg := &Config{
		Host:     host,
		Username: userName,
		Password: password,
		DbName:   dbName,
		Port:     port,
		SslMode:  sslMode,
		TimeZone: timeZone,
	}

	return cfg
}
