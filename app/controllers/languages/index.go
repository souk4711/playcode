package languages

import (
	"github.com/gin-gonic/gin"
)

type Language struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Langcode string `json:"langcode"`
}

func Index(c *gin.Context) {
	items := []Language{
		{Id: "c", Name: "C (GCC 10.2.1)", Langcode: "c"},
		{Id: "cpp", Name: "C++ (GCC 10.2.1)", Langcode: "cpp"},
		{Id: "d", Name: "D (2.100.0)", Langcode: "d"},
		{Id: "erlang", Name: "Erlang (25.0.4)", Langcode: "erlang"},
		{Id: "golang", Name: "Golang (1.19)", Langcode: "go"},
		{Id: "haskell", Name: "Haskell (GHC 9.4.2)", Langcode: "haskell"},
		{Id: "java", Name: "Java (OpenJDK 18)", Langcode: "java"},
		{Id: "javascript", Name: "Javascript (Node.js 18.8.0)", Langcode: "javascript"},
		{Id: "python", Name: "Python (3.10.6)", Langcode: "python"},
		{Id: "ruby", Name: "Ruby (3.1.2)", Langcode: "ruby"},
		{Id: "rust", Name: "Rust (1.63.0)", Langcode: "rust"},
		{Id: "scala", Name: "Scala (3.1.3)", Langcode: "scala"},
		{Id: "typescript", Name: "Typescript (4.8.2)", Langcode: "typescript"},
	}
	c.JSON(200, gin.H{
		"items": items,
	})
}
