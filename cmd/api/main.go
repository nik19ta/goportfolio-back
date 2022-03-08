package main

import (
	"go-just-portfolio/pkg/config"
	"go-just-portfolio/server"
	"log"
)

func main() {
	conf := config.GetConfig()

	app := server.NewApp()

	if err := app.Run(conf.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
