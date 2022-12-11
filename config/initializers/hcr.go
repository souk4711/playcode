package initializers

import (
	"os"

	"github.com/souk4711/playcode/pkg/hcr"
)

func InitializeHCR() error {
	uri := os.Getenv("CODERUNNER_URL")
	if client, err := hcr.NewClient(uri); err != nil {
		return err
	} else {
		hcr.API = client
		return nil
	}
}
