package runs

import (
	"github.com/gin-gonic/gin"
)

type RunResult struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func Create(c *gin.Context) {
	result := RunResult{
		Status: "ok",
		Reason: "",
		Stdout: "Hello, World",
		Stderr: "",
	}
	c.JSON(200, result)
}
