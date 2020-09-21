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

	err := newUser.Validate("create")
	if err != nil {
		ResERROR(w, http.StatusUnauthorized, err)
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
	err := user.Validate("login")
	if err != nil {
		ResERROR(w, http.StatusUnauthorized, err)
		return
	}

	info, err := u.db.GetUserByUserName(user.UserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ResERROR(w, http.StatusUnauthorized, errors.New("user not exists"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(user.Password))
	if err != nil {
		ResERROR(w, http.StatusUnauthorized, errors.New("Wrong Password or Wrong Username"))
		return
	}
	ResJSON(w, http.StatusOK, &Response{Data: info})
}
