package initializers

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/souk4711/playcode/config/utils"
)

func InitializeGin() error {
	env := os.Getenv("GIN_ENV")
	switch env {
	case utils.EnvDevelopment:
		gin.SetMode(gin.DebugMode)
	case utils.EnvTest:
		gin.SetMode(gin.TestMode)
	case utils.EnvProduction:
		gin.SetMode(gin.ReleaseMode)
	default:
		panic("unreachable!")
	}

	return nil
}
