package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api.port", "3333")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
	}

	return nil
}

func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetAPIConfig() string {
	return cfg.API.Port
}
