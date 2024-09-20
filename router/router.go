package router

import (
	"lf/gochat/internal/user"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter(userHandler *user.Handler) {
	router = gin.Default()

	router.POST("/signup", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	router.GET("/logout", userHandler.Logout)
}

func Start(address string) error {
	return router.Run(address)
}
