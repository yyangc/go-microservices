package handlers

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func (u *User) UserInfo(w http.ResponseWriter, r *http.Request) {
	id := getUserID(r)
	user, err := u.db.GetUserInfo(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ResERROR(w, http.StatusNotFound, errors.New("user not found"))
		return
	}
	if err != nil {
		ResERROR(w, http.StatusInternalServerError, err)
		return
	}
	ResJSON(w, http.StatusOK, &Response{Data: user})
}

func (u *User) UserOrderList(w http.ResponseWriter, r *http.Request) {
	id := getUserID(r)
	list, err := u.db.GetUserOrderList(id)
	if err != nil {
		ResERROR(w, http.StatusInternalServerError, err)
	}
	ResJSON(w, http.StatusOK, &Response{Data: list})
}
