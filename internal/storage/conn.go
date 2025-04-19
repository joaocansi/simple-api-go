package storage

import (
	"fmt"

	"github.com/joaocansi/simple-api/internal/config"
	storage "github.com/joaocansi/simple-api/internal/storage/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&storage.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
