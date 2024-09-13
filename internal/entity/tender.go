package entity

import "time"

type Tender struct {
	ID              int       `json:"-"`
	Name            string    `json:"name" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	ServiceType     string    `json:"serviceType" binding:"required"`
	OrganizationID  string    `json:"organizationId" binding:"required"`
	CreatorUsername string    `json:"creatorUsername" binding:"required"`
	Status          string    `json:"-"`
	Version         int       `json:"-"`
	CreatedAt       time.Time `json:"-"`
}
