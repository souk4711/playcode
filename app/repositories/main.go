package repositories

import (
	"github.com/uptrace/bun"

	"github.com/souk4711/playcode/app/repositories/languages"
)

var Language *languages.Repository

func Initialize(db *bun.DB) {
	Language = languages.NewRepository(db)
}
