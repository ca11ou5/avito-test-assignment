package tender_repo

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	pg *sqlx.DB
}

func New(pg *sqlx.DB) *Repository {
	return &Repository{
		pg: pg,
	}
}

func (r *Repository) SaveTender(tender *entity.Tender) (*entity.Tender, error) {}
