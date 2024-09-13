package main

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/controller/http"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/repository"
	tenderRepo "git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/repository/tender"
	tenderSvc "git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/service/tender"

	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/config"
)

func main() {
	var cfg config.Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		slog.Error("failed to read environment variables", "error", err)
	}

	db := repository.GetDBInstance(&cfg)

	tenderRepository := tenderRepo.New(db)
	tenderService := tenderSvc.New(tenderRepository)

	controller := http.NewController(tenderService, nil)

	err = controller.StartHTTPListening(&cfg)
	if err != nil {
		slog.Error("failed to start http listening", "error", err)
	}
}
