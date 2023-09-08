package database

import (
	"fmt"

	"github.com/brunoeduardodev/trellogo/internal/infra/config"
	"github.com/brunoeduardodev/trellogo/internal/infra/database/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDatabase() *gorm.DB {
	config := config.NewConfig().Database

	DSN := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", config.Host, config.User, config.Password, config.Name, config.Port)

	fmt.Printf("Will try to connect to DSN: %v\n", DSN)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	db.AutoMigrate(&users.UserModel{})

	if err != nil {
		panic("Couldn't connect to database")
	}

	return db
}
