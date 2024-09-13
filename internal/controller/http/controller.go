package http

import "git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/entity"

type TenderService interface {
	CreateTender(tender *entity.Tender) (*entity.Tender, error)
}

type BidService interface {
}

type Controller struct {
	TenderService
	BidService
}

func NewController(tenderSvc TenderService, bidSvc BidService) *Controller {
	return &Controller{
		TenderService: tenderSvc,
		BidService:    bidSvc,
	}
}
