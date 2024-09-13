package repository

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/config"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/entity"
	bid "git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/repository/bids"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/repository/tender"
)

type TenderClient interface {
	InsertTender(*entity.Tender) (*entity.Tender, error)
}

type BidClient interface {
}

type Repository struct {
	TenderClient
	BidClient
}

func New(cfg *config.Config) *Repository {
	db := connectToPostgres(cfg)

	return &Repository{
		TenderClient: tender.NewClient(db),
		BidClient:    bid.NewClient(db),
	}
}

func (r *Repository) InsertTender(tender *entity.Tender) (*entity.Tender, error) {
	return r.TenderClient.InsertTender(tender)
}
