module github.com/f1renze/the-architect

go 1.13

replace github.com/f1renze/the-architect => ./

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/jmoiron/sqlx v1.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro v1.18.0
	github.com/pquerna/otp v1.2.0
	github.com/tencentcloud/tencentcloud-sdk-go v3.0.121+incompatible
	go.uber.org/zap v1.12.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	google.golang.org/grpc v1.25.1
)
