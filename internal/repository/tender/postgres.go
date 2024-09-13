package tender

import (
	"github.com/jmoiron/sqlx"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/entity"
)

type Client struct {
	pg *sqlx.DB
}

func NewClient(db *sqlx.DB) *Client {
	return &Client{pg: db}
}

func (c *Client) InsertTender(tender *entity.Tender) (*entity.Tender, error) {
	return nil, nil
}
