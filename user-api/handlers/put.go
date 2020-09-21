package handlers

import (
	"errors"
	"github.com/yyangc/go-microservices/user-api/data"
	"net/http"
	"strconv"
)

func (u *User) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.FormValue("id"), 10, 64)
	if id == 0 || err != nil {
		ResERROR(w, http.StatusUnauthorized, errors.New("invalid user id"))
		return
	}
	sex, _ := strconv.ParseUint(r.FormValue("sex"), 10, 8)
	user := &data.User{
		ID:       id,
		UserName: r.FormValue("username"),
		Sex:      uint8(sex),
		Mail:     r.FormValue("mail"),
	}
	err = user.Validate("update")
	if err != nil {
		ResERROR(w, http.StatusUnauthorized, err)
		return
	}
	err = u.db.UpdateInfo(user)
	if err != nil {
		ResERROR(w, http.StatusInternalServerError, err)
		return
	}
	ResJSON(w, http.StatusOK, &Response{Message: "Update success"})
}
