package config

import (
	"github.com/gin-gonic/gin"

	"github.com/souk4711/playcode/app/controllers/languages"
	"github.com/souk4711/playcode/app/controllers/runs"
)

func routes(r *gin.Engine) {
	r.GET("/languages", languages.Index)
	r.POST("/runs", runs.Create)
}
