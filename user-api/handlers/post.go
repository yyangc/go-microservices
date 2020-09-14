package handlers

import (
	"database/sql"
	"github.com/yyangc/go-microservices/user-api/data"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	sex, _ := strconv.ParseUint(r.FormValue("sex"), 10, 8)
	newUser := &data.User{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
		Sex: uint8(sex),
		Mail: r.FormValue("mail"),
	}

	_, err := u.db.GetUserByUserName(newUser.UserName)
	if err != sql.ErrNoRows {
		ResJSON(w, http.StatusUnauthorized, &Response{Message: "username already exists"})
		return
	}

	id, err := u.db.CreateUser(newUser)
	if err != nil {
		ResERROR(w, http.StatusInternalServerError, err)
		return
	}
	info := map[string]int64{
		"id": id,
	}

	ResJSON(w, http.StatusOK, &Response{Data: info})
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	user := data.User{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	info, err := u.db.GetUserByUserName(user.UserName)
	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(user.Password))
	if err != nil {
		ResJSON(w, http.StatusUnauthorized, &Response{Message: "Wrong Password or Wrong Username"})
		return
	}
	ResJSON(w, http.StatusOK, &Response{Data: info})
}
