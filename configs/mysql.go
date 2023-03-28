package configs

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
)

// MySQLConfig define mysql configuration environment.
type MySQLConfig struct {
	DBMySQLHost     string `env:"DB_MYSQL_HOST" default:"localhost"`
	DBMySQLName     string `env:"DB_MYSQL_NAME" default:"example"`
	DBMySQLPort     int    `env:"DB_MYSQL_PORT" default:"3306"`
	DBMySQLUser     string `env:"DB_MYSQL_USER" default:"root"`
	DBMySQLPassword string `env:"DB_MYSQL_PASSWORD" default:"password"`
}

// DNS return connect string.
func (c *MySQLConfig) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBMySQLUser, c.DBMySQLPassword, c.DBMySQLHost, c.DBMySQLPort, c.DBMySQLName)
}

// LoadMySQLConfig load mysql config from env.
func LoadMySQLConfig() *MySQLConfig {
	cfg := &MySQLConfig{}
	mustParse(cfg)
	return cfg
}

// mustParse parse config from env using kong.
// kong: github.com/alecthomas/kong
func mustParse(cfg interface{}) {
	parser, _ := kong.New(cfg)
	_, err := parser.Parse([]string{})
	if err != nil {
		log.Fatal(err)
	}
}
