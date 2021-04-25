module github.com/jaswanth-gorripati/PGK/s2_Auth

go 1.15

replace github.com/jaswanth-gorripati/PGK/s2_Auth/configuration => ./configuration

replace github.com/jaswanth-gorripati/PGK/s2_Auth/controller => ./controller

replace github.com/jaswanth-gorripati/PGK/s2_Auth/dto => ./dto

replace github.com/jaswanth-gorripati/PGK/s2_Auth/models => ./models

replace github.com/jaswanth-gorripati/PGK/s2_Auth/routes => ./routes

replace github.com/jaswanth-gorripati/PGK/s2_Auth/services => ./services

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/githubnemo/CompileDaemon v1.2.1 // indirect
	github.com/go-playground/validator v9.31.0+incompatible // indirect
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/jaswanth-gorripati/PGK/s2_Auth/configuration v0.0.0-00010101000000-000000000000
	github.com/jaswanth-gorripati/PGK/s2_Auth/controller v0.0.0-00010101000000-000000000000 // indirect
	github.com/jaswanth-gorripati/PGK/s2_Auth/dto v0.0.0-00010101000000-000000000000
	github.com/jaswanth-gorripati/PGK/s2_Auth/models v0.0.0-00010101000000-000000000000
	github.com/jaswanth-gorripati/PGK/s2_Auth/routes v0.0.0-00010101000000-000000000000
	github.com/jaswanth-gorripati/PGK/s2_Auth/services v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.4.0
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2
)
