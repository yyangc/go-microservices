package handlers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/yyangc/go-microservices/user-api/data"
	"net/http"
	"strconv"
)

type User struct {
	l  *log.Logger
	db *data.UserDB
}

func NewUser(l *log.Logger, db *data.UserDB) *User {
	return &User{l: l, db: db}
}

func getUserID(r *http.Request) uint64 {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		// should never happen
		panic(err)
	}

	return uint64(id)
}
