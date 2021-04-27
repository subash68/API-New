module github.com/jaswanth-gorripati/PGK/s8_Notifications

go 1.15

replace github.com/jaswanth-gorripati/PGK/s8_Notifications/configuration => ./configuration

replace github.com/jaswanth-gorripati/PGK/s8_Notifications/controllers => ./controllers

replace github.com/jaswanth-gorripati/PGK/s8_Notifications/models => ./models

replace github.com/jaswanth-gorripati/PGK/s8_Notifications/routes => ./routes

replace github.com/jaswanth-gorripati/PGK/s8_Notifications/middleware => ./middleware

require (
	github.com/bwplotka/bingo v0.3.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/efficientgo/tools/core v0.0.0-20210122140009-1d4f98713811 // indirect
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jaswanth-gorripati/PGK/s8_Notifications/configuration v0.0.0-00010101000000-000000000000
	github.com/jaswanth-gorripati/PGK/s8_Notifications/controllers v0.0.0-00010101000000-000000000000 // indirect
	github.com/jaswanth-gorripati/PGK/s8_Notifications/middleware v0.0.0-00010101000000-000000000000 // indirect
	github.com/jaswanth-gorripati/PGK/s8_Notifications/models v0.0.0-00010101000000-000000000000
	github.com/jaswanth-gorripati/PGK/s8_Notifications/routes v0.0.0-00010101000000-000000000000
	github.com/razorpay/razorpay-go v0.0.0-20201204135735-096d3be7d2df // indirect
	github.com/stretchr/testify v1.6.1
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3 // indirect
	golang.org/x/tools/gopls v0.6.4 // indirect
	google.golang.org/api v0.36.0 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)
