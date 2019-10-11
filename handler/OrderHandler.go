package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wahyuhadi/gin-server/utils"
)

type OrderHandler struct {
	Log     *logrus.Logger
	Environ utils.Environ
}

func (orderHandler *OrderHandler) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ini orders handler",
	})
}

func NewOrdersHandler(logger *logrus.Logger, environ utils.Environ) *OrderHandler {
	return &OrderHandler{Log: logger, Environ: environ}
}

func (orderHandler *OrderHandler) Setup(r *gin.Engine) {
	r.GET("/orders", orderHandler.GetOne)
}
