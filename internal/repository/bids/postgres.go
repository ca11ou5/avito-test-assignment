package bid

import "github.com/jmoiron/sqlx"

type Client struct {
	pg *sqlx.DB
}

func NewClient(db *sqlx.DB) *Client {
	return &Client{pg: db}
}
