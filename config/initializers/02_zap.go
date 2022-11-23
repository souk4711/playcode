package initializers

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/souk4711/playcode/config/utils"
)

func InitializeZap() error {
	var config zap.Config

	env := os.Getenv("GIN_ENV")
	switch env {
	case utils.EnvDevelopment:
		fallthrough
	case utils.EnvTest:
		config = zap.NewDevelopmentConfig()
		config.OutputPaths = []string{"stdout", fmt.Sprintf("log/%s.log", env)}
	case utils.EnvProduction:
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{"stdout", fmt.Sprintf("log/%s.log", env)}
	default:
		panic("unreachable!")
	}

	if logger, err := config.Build(); err != nil {
		return fmt.Errorf("InitializeZap: %s", err.Error())
	} else {
		zap.ReplaceGlobals(logger)
		return nil
	}
}
