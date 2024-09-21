package router

import (
	"lf/gochat/internal/user"
	"lf/gochat/internal/webskt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter(userHandler *user.Handler, webSktHandler *webskt.Handler) {
	router = gin.Default()

	router.POST("/signup", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	router.GET("/logout", userHandler.Logout)

	router.POST("/rooms/create", webSktHandler.CreateRoom)
	router.GET("/rooms/join/:roomId", webSktHandler.JoinRoom)
	router.GET("/rooms", webSktHandler.GetRooms)
	router.GET("/rooms/clients/:roomId", webSktHandler.GetRoomClients)
}

func Start(address string) error {
	return router.Run(address)
}
