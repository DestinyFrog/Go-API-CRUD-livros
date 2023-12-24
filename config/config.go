package config

import "github.com/spf13/viper"

type apiconfig struct {
	Port string
}

type dbconfig struct {
	Name string
	Port string
	Host string
	User string
	Password string
}

var ApiConfig apiconfig
var DbConfig dbconfig

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	ApiConfig.Port = viper.GetString("API_PORT")

	DbConfig.Host = viper.GetString("DB_HOST")
	DbConfig.Port = viper.GetString("DB_PORT")
	DbConfig.Name = viper.GetString("DB_NAME")
	DbConfig.User = viper.GetString("DB_USER")
	DbConfig.Password = viper.GetString("DB_PASSWORD")
}