package configs

// ManticoreConfig ...
type ManticoreConfig struct {
	DBManticoreHost string `env:"DB_MANTICORE_HOST" default:"localhost"`
	DBManticorePort uint16 `env:"DB_MANTICORE_PORT" default:"9312"`
}

// LoadManticoreConfig load manticore search configuration from env.
func LoadManticoreConfig() *ManticoreConfig {
	cfg := &ManticoreConfig{}
	mustParse(cfg)
	return cfg
}
