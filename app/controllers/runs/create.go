package runs

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/souk4711/playcode/app/entities"
)

type CreateResponse struct {
	entities.RunResult
}

func NewCreateResponse() *CreateResponse {
	result := entities.RunResult{
		Status: "ok",
		Reason: "",
		Stdout: "Hello, World",
		Stderr: "",
	}
	return &CreateResponse{RunResult: result}
}

func Create(c *gin.Context) {
	c.JSON(http.StatusCreated, NewCreateResponse())
}
