package languages

import (
	"github.com/souk4711/playcode/app/types/models"
)

type IndexResponse struct {
	Items []models.Language `json:"items"`
}
