package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	pb "github.com/yyangc/go-microservices/order/protos/order"
	"github.com/yyangc/go-microservices/user-api/config"
	"github.com/yyangc/go-microservices/user-api/data"
	"github.com/yyangc/go-microservices/user-api/handlers"
	"google.golang.org/grpc"
)

func main() {
	l := initLogger()

	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	oClient := pb.NewOrderClient(conn)

	uDB, err := data.NewUserDB(l, oClient)
	if err != nil {
		l.Fatal(err)
	}

	//create the handlers
	u := handlers.NewUser(l, uDB)

	// create a new server
	r := mux.NewRouter()

	// handlers for API
	getR := r.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/user/info/{id:[0-9]+}", u.UserInfo)
	getR.HandleFunc("/user/order-list/{id:[0-9]+}", u.UserOrderList)

	postR := r.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/user", u.CreateUser)
	postR.HandleFunc("/login", u.Login)

	//putR := r.Methods(http.MethodPut).Subrouter()
	//putR.HandleFunc("/user/info/{id:[0-9]+}", u.UpdateUser)

	//deleteR := r.Methods(http.MethodDelete).Subrouter()
	//deleteR.HandleFunc("/user/{id:[0-9]+}", u.DeleteUser)

	srv := &http.Server{
		Addr:         ":" + config.Env.Port,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			l.Error(err.Error())
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan
	l.Info("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := srv.Shutdown(tc); err != nil {
		l.Warning("HTTP server Shutdown: %v", err.Error())
	}
}

func initLogger() *log.Logger {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{})
	l.SetLevel(log.InfoLevel)
	l.SetReportCaller(true)
	file, err := os.OpenFile("/Users/yang.c/go/src/github.com/yyangc/go-microservices/user-api/log/user.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Println(err.Error())
	}
	mw := io.MultiWriter(os.Stdout, file)
	l.SetOutput(mw)
	return l
}
