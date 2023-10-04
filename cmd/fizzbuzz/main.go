package main

import (
	"log"

	"github.com/felippegaede/fizz-buzz/api"
	"github.com/felippegaede/fizz-buzz/config"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server := api.NewServer()

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
