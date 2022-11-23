package config

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/souk4711/playcode/config/initializers"
	"github.com/souk4711/playcode/config/middlewares"
	"github.com/souk4711/playcode/config/utils"
)

var (
	Router *gin.Engine
)

func Initialize() {
	utils.LogFatalErr(initializers.InitializeDotenv())
	utils.LogFatalErr(initializers.InitializeGin())
	utils.LogFatalErr(initializers.InitializeZap())

	Router = newRouter()
}

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.RequestID())
	router.Use(middlewares.Ginzap(zap.L()))
	router.Use(middlewares.RecoveryWithZap(zap.L()))
	routes(router)
	return router
}
