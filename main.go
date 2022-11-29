package main

import (
	"github.com/souk4711/playcode/app/repositories"
	"github.com/souk4711/playcode/config"
)

func main() {
	config.Initialize()

	repositories.Initialize(config.DB)
	_ = config.Router.Run()
}
