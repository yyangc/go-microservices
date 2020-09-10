package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yyangc/go-microservices/order/data"
	pb "github.com/yyangc/go-microservices/order/protos/order"
	"github.com/yyangc/go-microservices/order/server"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
)

func main() {
	l := initLogger()

	db, err := data.NewOrderDB()
	if err != nil {
		l.Fatal(err)
	}
	o := server.NewOrderServer(l, db)

	gs := grpc.NewServer()
	pb.RegisterOrderServer(gs, o)
	listen, err := net.Listen("tcp", ":9092")
	if err != nil {
		l.Fatal(err.Error())
	}
	gs.Serve(listen)
}

func initLogger() *log.Logger {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{})
	l.SetLevel(log.InfoLevel)
	l.SetReportCaller(true)
	file, err := os.OpenFile("/Users/yang.c/go/src/github.com/yyangc/go-microservices/order/log/order.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Println(err.Error())
	}
	mw := io.MultiWriter(os.Stdout, file)
	l.SetOutput(mw)
	return l
}
