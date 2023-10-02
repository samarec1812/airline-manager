package main

import (
	"github.com/samarec1812/airline-manager/internal/app"
	"github.com/samarec1812/airline-manager/internal/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
