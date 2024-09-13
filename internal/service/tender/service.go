package tender

import "git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/entity"

type Repository interface {
	insertTender(*entity.Tender) (*entity.Tender, error)
}

type service struct{}
