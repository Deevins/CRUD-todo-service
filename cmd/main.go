package main

import (
	todoserver "github.com/deevins/todo-restAPI"
	"github.com/deevins/todo-restAPI/pkg/handler"
	"log"
)

func main() {
	srv := new(todoserver.Server)
	handlers := new(handler.Handler)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}
