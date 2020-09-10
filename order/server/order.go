package server

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	pb "github.com/yyangc/go-microservices/order/protos/order"
)

type OrderServer struct {
	l  *log.Logger
	db *sql.DB
}

func NewOrderServer(l *log.Logger, db *sql.DB) *OrderServer {
	return &OrderServer{l,  db}
}

func (o *OrderServer) GetUserOrderList(ctx context.Context, req *pb.UserRequest) (*pb.OrdersResponse, error) {
	id := req.UserId
	rows, err := o.db.Query("SELECT id, u_id, price, status, create_dt FROM orders WHERE u_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := make([]*pb.OrdersResponse_Order, 0)
	for rows.Next() {
		order := &pb.OrdersResponse_Order{}
		err = rows.Scan(&order.OrderId, &order.UserId,  &order.Price, &order.Status, &order.CreateDt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if rows.Err() != nil {
		return nil, err
	}
	res := pb.OrdersResponse{
		Orders: orders,
	}
	return &res, nil
}
