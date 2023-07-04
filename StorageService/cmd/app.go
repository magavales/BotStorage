package main

import (
	"StorageService/pkg/handler"
	"StorageService/pkg/server"
	"log"
)

type App struct {
}

func (a App) Do() {
	router := new(handler.Handler).InitRouter()

	serv := new(server.Server)
	err := serv.InitServer("8080", router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}
