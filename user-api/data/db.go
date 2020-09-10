package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/yyangc/go-microservices/order/protos/order"
	"github.com/yyangc/go-microservices/user-api/config"
)

type UserDB struct {
	l  *log.Logger
	db *sql.DB
	oc order.OrderClient
}

func NewUserDB(l *log.Logger, oc order.OrderClient) (*UserDB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.Env.MySQL.User,
		config.Env.MySQL.Password,
		config.Env.MySQL.Host,
		config.Env.MySQL.Port,
		config.Env.MySQL.Name)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return &UserDB{l, db, oc}, nil
}
