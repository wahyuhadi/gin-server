package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wahyuhadi/gin-server/utils"
)

type UserHandler struct {
	Log     *logrus.Logger
	Environ utils.Environ
}

func (userHandler *UserHandler) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ini User handler",
	})
}

func NewUsersHandler(logger *logrus.Logger, environ utils.Environ) *UserHandler {
	return &UserHandler{Log: logger, Environ: environ}
}

func (userHandler *UserHandler) Setup(r *gin.Engine) {
	r.GET("/users", userHandler.GetOne)
}
