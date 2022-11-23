package main

import (
	"github.com/souk4711/playcode/config"
)

func main() {
	config.Initialize()
	_ = config.Router.Run()
}
