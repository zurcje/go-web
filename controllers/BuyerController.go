package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	CardNumber int64  `json: "card_number"`
	FirstName  string `json:"first_name"`
	LasttName  string `json:"last_name"`
}

type Buyer struct {
	service buyers.Service
}

func NewBuyer(b buyers.Service) *Buyer {
	return &Buyer{
		service: NewBuyer,
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

		b, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, b)

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
		b, err := c.service.Store(req.CardNumber, req.FirstName, req.LasttName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, b)
	}
}
