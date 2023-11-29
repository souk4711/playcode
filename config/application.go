package config

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/zap"

	"github.com/souk4711/playcode/config/initializers"
	"github.com/souk4711/playcode/config/middlewares"
	"github.com/souk4711/playcode/config/utils"
)

var (
	Router *gin.Engine
	DB     *bun.DB
)

func Initialize() {
	utils.LogFatalErr(initializers.InitializeDotenv())
	utils.LogFatalErr(initializers.InitializeGin())
	utils.LogFatalErr(initializers.InitializeZap())
	utils.LogFatalErr(initializers.InitializeHCR())

	Router = newRouter()
	DB = newDB()
}

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.RequestID())
	router.Use(middlewares.Ginzap(zap.L()))
	router.Use(middlewares.RecoveryWithZap(zap.L()))
	routes(router)
	return router
}

func newDB() *bun.DB {
	dsn := os.Getenv("PLAYCODE_DATABASE_URL")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return db
}
