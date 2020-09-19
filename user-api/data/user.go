package data

import (
	"context"
	"database/sql"
	pb "github.com/yyangc/go-microservices/order/protos/order"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var SexList = []uint8{0, 1, 2}

type User struct {
	ID        uint64         `gorm:"column:id" json:"id"`
	UserName  string         `gorm:"column:username" json:"username"`
	Password  string         `gorm:"column:password" json:"-"`
	Status    uint8          `gorm:"column:status;default:1" json:"status"`
	Sex       uint8          `gorm:"column:sex;default:0" json:"sex"`
	Mail      string         `gorm:"column:mail" json:"mail"`
	CreatedDt time.Time      `gorm:"column:created_dt;default:null" json:"-"`
	UpdatedDt sql.NullString `gorm:"column:updated_dt;default:null" json:"-"`
}

func (u *UserDB) GetUserInfo(id int64) (*User, error) {
	user := new(User)
	res := u.db.First(&user, id)

	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (u *UserDB) GetUserByUserName(s string) (*User, error) {
	user := new(User)
	res := u.db.Where("username = ?", s).First(&user)

	if res.Error != nil {
		return nil, res.Error
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

func (u *UserDB) CreateUser(us *User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(us.Password), 14)
	if err != nil {
		return err
	}
	us.Password = string(bytes)
	res := u.db.Create(&us)
	if res.Error != nil {
		u.l.Error(res.Error)
		return res.Error
	}
	return nil
}

func (u *UserDB) UpdateInfo(us *User) error {
	res := u.db.Model(us).Select("sex", "mail").Updates(*us)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
