package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wahyuhadi/gin-server/handler"
	"github.com/wahyuhadi/gin-server/utils"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	var log = logrus.New()
	utils.InitLogFile(log)
	environ := utils.InitEnv(log)
	r := gin.New()

	userHandler := handler.NewUsersHandler(log, environ)
	userHandler.Setup(r)

	orderHandler := handler.NewOrdersHandler(log, environ)
	orderHandler.Setup(r)
	graceful.Run(environ.AppHost, 10*time.Second, r)
}
