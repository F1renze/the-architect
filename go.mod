module github.com/f1renze/the-architect

go 1.13

replace github.com/f1renze/the-architect => ./

require (
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.12
	github.com/jmoiron/sqlx v1.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	go.uber.org/zap v1.12.0
	google.golang.org/grpc v1.25.1
)
