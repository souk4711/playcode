package languages

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/souk4711/playcode/app/types/controllers/languages"
	"github.com/souk4711/playcode/app/types/models"
)

func Index(c *gin.Context) {
	items := []models.Language{
		{ID: "c", Name: "C (GCC 10.2.1)", Langcode: "c"},
		{ID: "cpp", Name: "C++ (GCC 10.2.1)", Langcode: "cpp"},
		{ID: "d", Name: "D (2.100.0)", Langcode: "d"},
		{ID: "erlang", Name: "Erlang (25.0.4)", Langcode: "erlang"},
		{ID: "golang", Name: "Golang (1.19)", Langcode: "go"},
		{ID: "haskell", Name: "Haskell (GHC 9.4.2)", Langcode: "haskell"},
		{ID: "java", Name: "Java (OpenJDK 18)", Langcode: "java"},
		{ID: "javascript", Name: "Javascript (Node.js 18.8.0)", Langcode: "javascript"},
		{ID: "python", Name: "Python (3.10.6)", Langcode: "python"},
		{ID: "ruby", Name: "Ruby (3.1.2)", Langcode: "ruby"},
		{ID: "rust", Name: "Rust (1.63.0)", Langcode: "rust"},
		{ID: "scala", Name: "Scala (3.1.3)", Langcode: "scala"},
		{ID: "typescript", Name: "Typescript (4.8.2)", Langcode: "typescript"},
	}
	c.JSON(http.StatusOK, languages.IndexResponse{
		Items: items,
	})
}
