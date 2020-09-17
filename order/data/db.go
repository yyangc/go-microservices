package data

import (
	"fmt"
	"github.com/yyangc/go-microservices/order/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewOrderDB() (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.PSQL.Host,
		config.Env.PSQL.Port,
		config.Env.PSQL.User,
		config.Env.PSQL.Password,
		config.Env.PSQL.Name)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
