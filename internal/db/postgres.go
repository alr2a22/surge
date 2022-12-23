package db

import (
	"fmt"
	"surge/internal/config"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var onceDB sync.Once

func GetDBConn() *gorm.DB {
	onceDB.Do(func() {
		var err error
		dsn := getDsn()
		logrus.Debugf("Database connetion string: %s\n", dsn)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.LogLevel(logrus.GetLevel() - 1))})
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Infoln("Connection established to postgres Database")
	})
	return db
}

func getDsn() string {
	cfg := config.GetConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
		cfg.PostgresPort)
}

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
