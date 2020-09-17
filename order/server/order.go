package server

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "github.com/yyangc/go-microservices/order/protos/order"
	"gorm.io/gorm"
)

type OrderServer struct {
	l  *log.Logger
	db *gorm.DB
}

func NewOrderServer(l *log.Logger, db *gorm.DB) *OrderServer {
	return &OrderServer{l, db}
}
func (o *OrderServer) GetUserOrderList(ctx context.Context, req *pb.UserRequest) (*pb.OrdersResponse, error) {
	id := req.UserId
	orders := make([]*pb.OrdersResponse_Order, 0)
	res := o.db.Table("orders").
		Select("id as order_id, u_id as user_id, price, status, to_char(created_dt, 'YYYY-MM-DD HH24:MI:SS') as created_dt").
		Where("u_id = ?", id).
		Find(&orders)

	if res.Error != nil {
		o.l.Info(res.Error)
		return nil, res.Error
	}

	data := pb.OrdersResponse{
		Orders: orders,
	}
	return &data, nil
}
