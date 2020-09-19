package handlers

import (
	"errors"
	"github.com/yyangc/go-microservices/user-api/data"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	sex, _ := strconv.ParseUint(r.FormValue("sex"), 10, 8)
	newUser := &data.User{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
		Sex:      uint8(sex),
		Mail:     r.FormValue("mail"),
	}

	if !newUser.CheckPassword() {
		ResERROR(w, http.StatusUnauthorized, errors.New("Create your password using 8 characters or more."))
		return
	}
	if !newUser.CheckUserName() {
		ResERROR(w, http.StatusUnauthorized, errors.New("Choose a username 6–30 characters long. Your username can be any combination of letters, numbers, or symbols."))
		return
	}
	if !newUser.CheckSex() {
		ResERROR(w, http.StatusUnauthorized, errors.New("invalid Sex"))
		return
	}
	if !newUser.CheckMail() {
		ResERROR(w, http.StatusUnauthorized, errors.New("invalid Mail"))
		return
	}

	// check if username exists
	user, err := u.db.GetUserByUserName(newUser.UserName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ResERROR(w, http.StatusInternalServerError, err)
		return
	}

	if user != nil {
		ResERROR(w, http.StatusUnauthorized, errors.New("username already exists"))
		return
	}

	// create user
	err = u.db.CreateUser(newUser)
	if err != nil {
		ResERROR(w, http.StatusInternalServerError, err)
		return
	}

	info := map[string]uint64{
		"id": newUser.ID,
	}
	ResJSON(w, http.StatusOK, &Response{Data: info})
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	user := data.User{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	info, err := u.db.GetUserByUserName(user.UserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ResJSON(w, http.StatusUnauthorized, &Response{Message: "user not exists"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(user.Password))
	if err != nil {
		ResJSON(w, http.StatusUnauthorized, &Response{Message: "Wrong Password or Wrong Username"})
		return
	}
	ResJSON(w, http.StatusOK, &Response{Data: info})
}
