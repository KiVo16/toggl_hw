package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	DefaultHost   = "0.0.0.0"
	DefaultPort   = "3000"
	DefaultDBFile = "test.db"
)

func bindEnv(inputs ...string) error {
	err := viper.BindEnv(inputs...)
	if err != nil {
		return fmt.Errorf("[config] bind env err: %v", err)
	}
	return nil
}

func LoadAndValidateConfig() (*Config, error) {

	configPath := flag.String("c", "", "config path")
	flag.Parse()

	viper.SetDefault("api.host", DefaultHost)
	viper.SetDefault("api.port", DefaultPort)
	viper.SetDefault("db.file", DefaultDBFile)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := bindEnv("api.port", "PORT")
	if err != nil {
		return nil, err
	}

	err = bindEnv("auth.jwtSecret", "AUTH_JWTSECRET")
	if err != nil {
		return nil, err
	}

	if configPath != nil && len(*configPath) > 0 {
		viper.SetConfigFile(*configPath)
	}

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("[config] read config err: %v", err)
		}
	}

	conf := Config{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("[config] unmarshal config err: %v", err)
	}

	err = conf.validate()
	if err != nil {
		return nil, fmt.Errorf("[config] validate config err: %v", err)
	}

	conf.prepare()

	return &conf, nil
}
