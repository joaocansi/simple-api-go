package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, engine *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
