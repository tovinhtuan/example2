package storages

import (
	"ex2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PSQLManager struct {
	*gorm.DB
}

func NewPSQLManager() (*PSQLManager, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=public port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}
	err1 := db.AutoMigrate(
		&models.User{},
		&models.AuthToken{},
	)
	if err1 != nil {
		return nil, err
	}
	return &PSQLManager{db}, nil
}