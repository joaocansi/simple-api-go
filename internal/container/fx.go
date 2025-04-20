package container

import (
	"net/http"

	"github.com/joaocansi/simple-api/internal/config"
	"github.com/joaocansi/simple-api/internal/server"
	"github.com/joaocansi/simple-api/internal/token"
	"github.com/joaocansi/simple-api/internal/users"
	"github.com/joaocansi/simple-api/storage"
	"go.uber.org/fx"
)

func Init() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			storage.NewConnection,
			users.NewUserHandler,
			users.NewUserService,
			token.NewTokenService,
			server.NewServerEngine,
			server.NewHTTPServer,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
