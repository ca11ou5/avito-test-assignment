package model

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/entity"
	"time"
)

func NewTenderEntity(req *NewTenderRequest) *entity.Tender {
	return &entity.Tender{
		Name:            req.Name,
		Description:     req.Description,
		ServiceType:     req.ServiceType,
		OrganizationID:  req.OrganizationID,
		CreatorUsername: req.OrganizationID,
		Status:          "CREATED",
		CreatedAt:       time.Now(),
	}
}
