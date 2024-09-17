package main

import (
	"fmt"
	"lf/gochat/db"
	"log"
)

func main() {
	_, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize database connection %s:", err)
	}

	fmt.Println("database connected")
}
