package main

import (
	"go-web/controllers"
	"go-web/internal/domains/buyer"
	"go-web/internal/domains/file"

	"github.com/gin-gonic/gin"
)

func main() {

	db := file.New(file.FileType, "../../buyer.json")

	repository := buyer.NewRepository(db)
	service := buyer.NewService(repository)
	controller := controllers.NewBuyers(service)

	r := gin.Default()
	buyer := r.Group("/buyers")
	buyer.POST("/", controller.Store())
	buyer.GET("/", controller.GetAll())

	r.Run()
}
