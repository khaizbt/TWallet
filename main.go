package main

import (
	"TWallet/auth"
	"TWallet/handler"
	"TWallet/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/twallet_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository) //parsing data input dari user(Struct input) ke Struct user
	authService := auth.NewServiceAuth()

	userHandler := handler.NewuserHandler(userService, authService) //Parsing Input dari user ke user Service

	router := gin.Default()
	api := router.Group("/api/v1/")
	api.POST("register", userHandler.RegisterUser)
	api.POST("login", userHandler.Login)
	api.POST("check-email", userHandler.CheckEmailAvailablity)

	router.Run()
}
