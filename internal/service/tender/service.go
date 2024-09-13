package tenderSvc

import "git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/entity"

type repository interface {
	SaveTender(*entity.Tender) (*entity.Tender, error)
}

type Service struct {
	repository
}

func New(tenderRepo repository) *Service {
	return &Service{
		repository: tenderRepo,
	}
}

func (s *Service) CreateTender(tender *entity.Tender) (*entity.Tender, error) {
	return s.repository.SaveTender(tender)
}
