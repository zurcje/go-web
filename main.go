package main

import (
	"go-web/controllers"
	"go-web/internal/domains/buyer"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := buyer.NewRepository()
	service := buyer.NewService(repository)
	controller := controllers.NewBuyers(service)

	r := gin.Default()
	buyer := r.Group("/buyers")
	buyer.POST("/", controller.Store())
	buyer.GET("/", controller.GetAll())

	r.Run()
}
