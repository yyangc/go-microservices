package data

import (
	"context"
	"database/sql"
	"errors"
	pb "github.com/yyangc/go-microservices/order/protos/order"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var SexList = []uint8{0, 1, 2}

type User struct {
	ID       uint64         `json:"id"`
	UserName string         `json:"username"`
	Password string			`json:"-"`
	Status   uint8          `json:"status"`
	Sex      uint8          `json:"sex"`
	Mail     string         `json:"mail"`
	CreateDt sql.NullString `json:"-"`
	UpdateDt sql.NullString `json:"-"`
}

func (u *UserDB) GetUserInfo(id int64) (*User, error) {
	user := new(User)
	err := u.db.QueryRow("SELECT * FROM user WHERE id = ?", id).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
		&user.Status,
		&user.Sex,
		&user.Mail,
		&user.CreateDt,
		&user.UpdateDt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDB) GetUserByUserName(s string) (*User, error) {
	user := new(User)
	err := u.db.QueryRow("SELECT * FROM user WHERE username = ?", s).Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
		&user.Status,
		&user.Sex,
		&user.Mail,
		&user.CreateDt,
		&user.UpdateDt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDB) GetUserOrderList(id int64) (*pb.OrdersResponse, error) {
	req := &pb.UserRequest{
		UserId: id,
	}
	list, err := u.oc.GetUserOrderList(context.Background(), req)
	if err != nil {
		u.l.Error(err)
	}
	return list, err
}

func (u *UserDB) CreateUser(us *User) (int64, error) {
	if us.UserName == "" || len(us.UserName) > 20 {
		return 0, errors.New("invalid username")
	}
	if us.Password == "" || len(us.Password) < 8 {
		return 0, errors.New("invalid password")
	}
	if isSex := checkSex(&us.Sex); isSex != true{
		return 0, errors.New("invalid Sex")
	}
	if isEmail := isEmailValid(&us.Mail); isEmail != true {
		return 0, errors.New("invalid Mail")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(us.Password), 14)
	if err != nil {
		return 0, err
	}
	us.Password = string(bytes)
	result, error := u.db.Exec("INSERT INTO user (username, password, sex, mail, create_dt) VALUES (?, ?, ?, ?, ?)",
		us.UserName,
		us.Password,
		us.Sex,
		us.Mail,
		time.Now().Format("2006-01-02 15:04:05"),
		)
	if error != nil {
		u.l.Error(error)
		return 0, error
	}
	id, _ := result.LastInsertId()
	return id, nil
}
