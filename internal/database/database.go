package database

import (
	"fmt"

	"github.com/Venukishore-R/CODECRAFT_BW_02/config"
	"github.com/Venukishore-R/CODECRAFT_BW_02/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", config.DBHOST, config.DBPORT, config.DBUSERNAME, config.DBPASSWORD, config.DBNAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
