package databases

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/manticoresoftware/go-sdk/manticore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"mysql-manticore-example/configs"
)

var onceMySQL sync.Once
var onceManticore sync.Once
var mysqlDB *gorm.DB
var manticoreClient = &manticore.Client{}

// InitializeMySQL get db connection instance.
func InitializeMySQL(config configs.MySQLConfig) *gorm.DB {
	onceMySQL.Do(func() {
		var err error
		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,         // Disable color
			},
		)
		mysqlDB, err = gorm.Open(mysql.Open(config.DNS()), &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			panic(fmt.Errorf("failed to connect mysql db: %+v", err))
		}
	})
	return mysqlDB
}

// InitializeManticore get manticore connection client.
func InitializeManticore(config configs.ManticoreConfig) *manticore.Client {
	onceManticore.Do(func() {
		client := manticore.NewClient()
		client.SetServer(config.DBManticoreHost, config.DBManticorePort)
		_, err := client.Open()
		if err != nil {
			panic(fmt.Errorf("failed to connect manticoresearch in mysql connection: %+v", err))
		}
		manticoreClient = &client
	})
	return manticoreClient
}
