module github.com/yyangc/go-microservices/user-api

go 1.14


require (
	github.com/yyangc/go-microservices/order v0.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.1
	google.golang.org/grpc v1.32.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/yyangc/go-microservices/order => ../order
