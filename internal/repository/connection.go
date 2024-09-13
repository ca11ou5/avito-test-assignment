package repository

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/config"
)

func connectToPostgres(cfg *config.Config) *sqlx.DB {
	dsn := createDSN(cfg)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		slog.Error("failed to connect postgres", "error", err)
	}

	if err = db.Ping(); err != nil {
		slog.Error("failed to ping postgres", "error", err)
	}

	slog.Info("successfully connected to postgres")
	return db
}

func createDSN(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.PostgresHost, cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresDatabase, cfg.PostgresPort)

	return dsn
}
