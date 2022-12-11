package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Language struct {
	bun.BaseModel          `bun:"table:languages"`
	ID                     uuid.UUID `bun:"id,pk,type:uuid"`
	CreatedAt              time.Time `bun:"created_at"`
	UpdatedAt              time.Time `bun:"updated_at"`
	Name                   string    `bun:"name"`
	Langcode               string    `bun:"langcode"`
	CodeRunnerLangcode     string    `bun:"cr_langcode"`
	CodeRunnerMainFilename string    `bun:"cr_main_filename"`
}
