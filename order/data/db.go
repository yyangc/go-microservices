package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/yyangc/go-microservices/order/config"
)


func NewOrderDB() (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.PSQL.Host,
		config.Env.PSQL.Port,
		config.Env.PSQL.User,
		config.Env.PSQL.Password,
		config.Env.PSQL.Name)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
