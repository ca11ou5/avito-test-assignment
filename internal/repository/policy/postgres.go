package policy

import "github.com/jmoiron/sqlx"

type Repository struct {
	pg *sqlx.DB
}

func NewRepository(pg *sqlx.DB) *Repository {
	return &Repository{
		pg: pg,
	}
}

const EmployerIDOnUsernameQuery = "SELECT id FROM employee WHERE username = $1"

func (r *Repository) CheckUserResponsible(username, organizationId string) error {

}
