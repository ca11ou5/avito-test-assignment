package bid

import "github.com/jmoiron/sqlx"

type Repository struct {
	pg *sqlx.DB
}
