package main

import (
	"github.com/go-api/app"
	"github.com/go-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8888")
}
