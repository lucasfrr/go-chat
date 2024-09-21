package main

import (
	"lf/gochat/db"
	"lf/gochat/internal/user"
	"lf/gochat/internal/webskt"
	"lf/gochat/router"
	"log"
)

func main() {
	db, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize database connection %s:", err)
	}

	userRepository := user.NewRepository(db.GetDb())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := webskt.NewHub()
	webSktHandler := webskt.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, webSktHandler)
	router.Start("0.0.0.0:8080")
}
