package languages

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/souk4711/playcode/app/models"
)

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Find(ctx context.Context, language *models.Language) error {
	return repo.db.NewSelect().Model(language).WherePK().Scan(ctx)
}

func (repo *Repository) FindAll(ctx context.Context, languages *[]models.Language) error {
	return repo.db.NewSelect().Model((*models.Language)(nil)).Scan(ctx, languages)
}
