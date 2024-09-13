package main

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/config"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725737622-team-78771/internal/controller/http"
)

func main() {
	var cfg config.Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		slog.Error("failed to read environment variables", "error", err)
	}

	// repo := repository.New(&cfg)
	// svc := service.New(repo)

	controller := http.NewController("asd")
	err = controller.StartHTTPListening(&cfg)
	if err != nil {
		slog.Error("failed to start http listening", "error", err)
	}
}
