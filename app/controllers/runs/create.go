package runs

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/souk4711/playcode/app/types/controllers/runs"
	"github.com/souk4711/playcode/app/types/models"
)

func Create(c *gin.Context) {
	result := models.RunResult{
		Status: "ok",
		Reason: "",
		Stdout: "Hello, World",
		Stderr: "",
	}
	c.JSON(http.StatusCreated, runs.CreateResponse{RunResult: result})
}
