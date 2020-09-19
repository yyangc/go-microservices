package data

import (
	"errors"
	"regexp"
	"strconv"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func GetUserId(s string) (uint64, error) {
	isInt, _ := regexp.MatchString("[0-9]+", s)
	if !isInt {
		return 0, errors.New("Invalid UserId")
	}

	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("Invalid UserId")
	}
	return id, nil
}

func (u *User) CheckSex() bool {
	for _, v := range SexList {
		if u.Sex == v {
			return true
		}
	}
	return false
}

func (u *User) CheckMail() bool {
	if len(u.Mail) < 3 && len(u.Mail) > 254 {
		return false
	}
	return emailRegex.MatchString(u.Mail)
}

func (u *User) CheckUserName() bool {
	if u.UserName == "" || len(u.UserName) < 6 || len(u.UserName) > 30 {
		return false
	}
	return true
}

func (u *User) CheckPassword() bool {
	if u.Password == "" || len(u.Password) < 8 {
		return false
	}
	return true
}
