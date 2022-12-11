package runs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/souk4711/playcode/app/entities"
	"github.com/souk4711/playcode/app/models"
	"github.com/souk4711/playcode/app/repositories"
	"github.com/souk4711/playcode/pkg/hcr"
	pb "github.com/souk4711/playcode/pkg/hcr/pb/runs"
)

type CreateRequest struct {
	LanguageID string           `json:"language_id" binding:"required"`
	RunFile    entities.RunFile `json:"file" binding:"required"`
}

type CreateResponse struct {
	entities.RunResult
}

func NewCreateResponse(r *pb.CreateResponse) *CreateResponse {
	result := entities.RunResult{
		Status: r.Status,
		Reason: r.Reason,
		Stdout: r.Stdout,
		Stderr: r.Stderr,
	}
	return &CreateResponse{RunResult: result}
}

func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var language = models.Language{ID: uuid.MustParse(req.LanguageID)}
	if err := repositories.Language.Find(c, &language); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	lang := language.CodeRunnerLangcode
	file := pb.File{Name: language.CodeRunnerMainFilename, Content: req.RunFile.Content}
	r := pb.CreateRequest{Lang: lang, Files: []*pb.File{&file}}

	if r, err := hcr.API.Runs.Create(c, &r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, NewCreateResponse(r))
	}
}
