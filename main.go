package main

import (
	"github.com/gin-gonic/gin"

	routes "github.com/souk4711/playcode/app/controllers"
)

func main() {
	app := gin.Default()

	routes.Configure(app)
	_ = app.Run()
}
