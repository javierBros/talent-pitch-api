package cmd

import (
	"project/application/app"
	"project/config"
)

func Execute() {
	config.LoadConfig()
	app.Start()
}
