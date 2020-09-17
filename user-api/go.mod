module github.com/yyangc/go-microservices/user-api

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.1
	github.com/yyangc/go-microservices/order v0.0.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0 // indirect
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.1
)

replace github.com/yyangc/go-microservices/order => ../order
