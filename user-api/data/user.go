package data

import (
	"context"
	"database/sql"
	pb "github.com/yyangc/go-microservices/order/protos/order"
)

type User struct {
	ID       uint64         `json:"id"`
	UserName string         `json:"username"`
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
