package cmd

import (
	"github.com/talent-pitch-api/application/app"
	"github.com/talent-pitch-api/config"
)

func Execute() {
	config.LoadConfig()
	app.Start()
}
