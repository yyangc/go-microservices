package data

import (
	"errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9]{6,30}$")

var ErrorUesrName = "Choose a username 6–30 characters long. Your username can be any combination of letters, numbers, or symbols."
var ErrorPassword = map[string]string{
	"len":    "Passwords must be at least 8 characters in length",
	"digit":  "Password must contain at least one number digit (ex: 0, 1, 2, 3, etc.)",
	"lower":  "Password must contain at least one lowercase letter.",
	"upper":  "Password must contain at least one uppercase, or capital, letter (ex: A, B, etc.)",
	"symbol": "Password must contain at least one special character -for example: $, #, @, !,%,^,&,*,(,)",
}

func (u *User) CheckSex() error {
	for _, v := range SexList {
		if u.Sex == v {
			return nil
		}
	}
	u.Sex = 0
	return nil
}

func (u *User) CheckMail() error {
	if u.Mail == "" {
		return errors.New("Empty Mail")
	}
	if len(u.Mail) < 3 && len(u.Mail) > 254 {
		return errors.New("Mail Invalid")
	}
	if !emailRegex.MatchString(u.Mail) {
		return errors.New("Mail Invalid")
	}
	return nil
}

func (u *User) CheckUserName() error {
	if !usernameRegex.MatchString(u.UserName) {
		return errors.New(ErrorUesrName)
	}
	return nil
}

func (u *User) CheckPassword() error {
	if len(u.Password) < 8 {
		return errors.New(ErrorPassword["len"])
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, u.Password); !b || err != nil {
		return errors.New(ErrorPassword["digit"])
	}
	if b, err := regexp.MatchString(a_z, u.Password); !b || err != nil {
		return errors.New(ErrorPassword["lower"])
	}
	if b, err := regexp.MatchString(A_Z, u.Password); !b || err != nil {
		return errors.New(ErrorPassword["upper"])
	}
	if b, err := regexp.MatchString(symbol, u.Password); !b || err != nil {
		return errors.New(ErrorPassword["symbol"])
	}
	return nil
}

func (u *User) Validate(action string) error {
	var s []func() error
	switch strings.ToLower(action) {
	case "create":
		s = []func() error{
			u.CheckUserName,
			u.CheckPassword,
			u.CheckSex,
			u.CheckMail,
		}
	case "update":
		s = []func() error{
			u.CheckSex,
			u.CheckMail,
		}
	case "login":
		if u.UserName == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
	}

	for _, v := range s {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}
