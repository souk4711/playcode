package initializers

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/souk4711/playcode/config/utils"
)

func InitializeDotenv() error {
	if env, ok := os.LookupEnv("GIN_ENV"); ok {
		switch env {
		case utils.EnvDevelopment:
		case utils.EnvTest:
		case utils.EnvProduction:
		default:
			return fmt.Errorf("InitializeDotenv: GIN_ENV unknown: %s", env)
		}
	} else {
		os.Setenv("GIN_ENV", utils.EnvDevelopment)
	}

	filepath := ".env"
	if _, err := os.Stat(filepath); !errors.Is(err, os.ErrNotExist) {
		if err := godotenv.Load(filepath); err != nil {
			return fmt.Errorf("InitializeDotenv: %s", err.Error())
		}
	}

	return nil
}
