package handlers

import (
	"database/sql"
	"net/http"
)

func (u *User) UserInfo(w http.ResponseWriter, r *http.Request) {
	id := getUserID(r)
	user, err := u.db.GetUserInfo(id)

	if err == sql.ErrNoRows {
		ResJSON(w, http.StatusNotFound, &Response{Message: "user not found"})
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
	if err!= nil {
		ResERROR(w, http.StatusInternalServerError, err)
	}
	ResJSON(w, http.StatusOK, &Response{Data: list})
}
