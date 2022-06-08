package main

import (
	"go-web/controllers"
	"go-web/internal/domains/buyer"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := buyer.NewRepository()
	service := buyer.NewService(repository)
	controller := controllers.NewBuyer(service)

	r := gin.Default()
	buyers := r.Group("/buyers")
	buyers.POST("/", controller.Store())
	buyers.GET("/", controller.GetAll())

	r.Run()
}
