package config

import (
	"fmt"

	"github.com/spf13/viper"
)


type Config struct{
	Server ServerConfig
	Postgres PostgresConfig
	JWTKey string
}


type ServerConfig struct{
	Host string
	Port string
}

type PostgresConfig struct{
	DbName string `yaml:"db_name"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Port string `yaml:"port"`
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error){
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct, %v", err)
	}
	return &c, nil
}