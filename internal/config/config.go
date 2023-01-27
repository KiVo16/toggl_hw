package config

import (
	"fmt"
)

type Mode = int

const (
	ModeHttp = iota
	ModeGrpc
)

type Config struct {
	API      APIConfig  `mapstructure:"api"`
	Auth     AuthConfig `mapstructure:"auth"`
	Database DBConfig   `mapstructure:"db"`
}

func (c Config) validate() error {
	err := c.Auth.validate()
	if err != nil {
		return err
	}

	err = c.API.validate()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) prepare() {
	c.API.prepare()
}

type APIConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	ModeString string `mapstructure:"mode"`
	Mode       Mode   `mapstructure:"-"`
}

func (c APIConfig) validate() error {
	if c.ModeString != "http" && c.ModeString != "grpc" {
		return fmt.Errorf("invalid mode - available 'grpc' or 'http'")
	}

	return nil
}

func (c *APIConfig) prepare() {
	if c.ModeString == "grpc" {
		c.Mode = ModeGrpc
		return
	}

	c.Mode = ModeHttp
}

type AuthConfig struct {
	JWTSecret string `mapstructure:"jwtSecret"`
}

func (c AuthConfig) validate() error {
	if len(c.JWTSecret) == 0 {
		return fmt.Errorf("jwtSecret cannot be empty")
	}
	return nil
}

type DBConfig struct {
	File string `mapstructure:"file"`
}
