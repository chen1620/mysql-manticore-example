package configs

import "github.com/alecthomas/kong"

// AppConfig define app configuration.
type AppConfig struct {
	HTTPAddress     string          `env:"HTTP_ADDRESS" default:":8080"`
	ManticoreConfig ManticoreConfig `kong:"embed"`
	MySQLConfig     MySQLConfig     `kong:"embed"`
}

// LoadAppConfig get app configuration from env.
func LoadAppConfig() *AppConfig {
	var cfg AppConfig
	kong.Parse(&cfg)
	return &cfg
}
