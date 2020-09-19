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
	mw := io.MultiWriter(os.Stdout)
	l.SetOutput(mw)
	return l
}
