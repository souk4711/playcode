package runs

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/souk4711/playcode/app/models"
)

func Create(c *gin.Context) {
	result := models.RunResult{
		Status: "ok",
		Reason: "",
		Stdout: "Hello, World",
		Stderr: "",
	}
	c.JSON(http.StatusCreated, result)
}
