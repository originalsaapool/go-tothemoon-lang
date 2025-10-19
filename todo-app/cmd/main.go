package main

import (
	"log"
	"todo-app"
	"todo-app/pkg/handler"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers:= handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running the http server %s", err.Error())
	}
}
