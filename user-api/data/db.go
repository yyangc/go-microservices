package data

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/yyangc/go-microservices/order/protos/order"
	"github.com/yyangc/go-microservices/user-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDB struct {
	l  *log.Logger
	db *gorm.DB
	oc order.OrderClient
}

func NewUserDB(l *log.Logger, oc order.OrderClient) (*UserDB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.Env.MySQL.User,
		config.Env.MySQL.Password,
		config.Env.MySQL.Host,
		config.Env.MySQL.Port,
		config.Env.MySQL.Name)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &UserDB{l, db, oc}, nil
}
