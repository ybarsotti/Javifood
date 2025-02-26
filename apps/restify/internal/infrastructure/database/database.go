package database

import (
	"fmt"
	"javifood-restify/config"
	"javifood-restify/internal/infrastructure/database/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase() error {
	dbDsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.EnvConfig.DB.Host,
		config.EnvConfig.DB.User,
		config.EnvConfig.DB.Password,
		config.EnvConfig.DB.DB,
		config.EnvConfig.DB.Port,
		config.EnvConfig.DB.SslMode,
		config.EnvConfig.DB.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("connected to database")
	err = db.AutoMigrate(&model.Restaurant{})
	if err != nil {
		log.Error("error in migration. ", err.Error())
	}
	DBConn = db
	return nil
}
