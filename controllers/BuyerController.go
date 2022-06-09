package controllers

import (
	"go-web/internal/domains/buyer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	CardNumber int64  `json: "cardNumber"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
}

type Buyer struct {
	service buyer.Service
}

func NewBuyers(newBuyer buyer.Service) *Buyer {
	return &Buyer{
		service: newBuyer,
	}
}

func (c *Buyer) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "1234567" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		newBuyer, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, newBuyer)

	}
}

func (c *Buyer) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "1234567" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return

		}
		newBuyer, err := c.service.Store(req.CardNumber, req.FirstName, req.LastName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, newBuyer)
	}
}
