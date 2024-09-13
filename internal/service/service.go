package service

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/service/bids"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/service/tender"
)

type repository interface {
	bids.Repository
	tender.Repository
}

type Service struct {
	repo repository
}

func New(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}
