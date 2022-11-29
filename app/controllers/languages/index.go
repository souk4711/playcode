package languages

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/souk4711/playcode/app/entities"
	"github.com/souk4711/playcode/app/models"
	"github.com/souk4711/playcode/app/repositories"
)

type IndexResponse struct {
	Items []entities.Language `json:"items"`
}

func NewIndexResponse(languages []models.Language) *IndexResponse {
	var items []entities.Language
	for _, language := range languages {
		items = append(items, entities.Language{
			ID:       language.ID.String(),
			Name:     language.Name,
			Langcode: language.Langcode,
		})
	}
	return &IndexResponse{Items: items}
}

func Index(c *gin.Context) {
	var languages []models.Language
	if err := repositories.Language.FindAll(c, &languages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, NewIndexResponse(languages))
}
