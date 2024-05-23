package api

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"sync"
	"time"

	"github.com/bahattincinic/fitwave"
	"github.com/bahattincinic/fitwave/config"
	"github.com/bahattincinic/fitwave/database"
	"github.com/bahattincinic/fitwave/importer"
	"github.com/bahattincinic/fitwave/queue"
	"github.com/bahattincinic/fitwave/strava"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type API struct {
	ec  *echo.Echo
	ctx context.Context
	cfg *config.Config
	db  *database.Database
	log *zap.Logger
	st  *strava.Strava
	im  *importer.Importer
	q   *queue.Queue
}

var uiFS fs.FS

func RunAPI(ctx context.Context, wg *sync.WaitGroup, log *zap.Logger, db *database.Database, cfg *config.Config, st *strava.Strava, im *importer.Importer, q *queue.Queue) {
	srv := &API{
		ec:  echo.New(),
		ctx: ctx,
		cfg: cfg,
		db:  db,
		log: log,
		st:  st,
		im:  im,
		q:   q,
	}
	srv.ec.Server.IdleTimeout = 120 * time.Second

	srv.ec.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	srv.setupHandlers()
	srv.setupSwagger()
	srv.setupCors()

	{
		var err error
		uiFS, err = fs.Sub(fitwave.UI, "ui/dist")
		if err != nil {
			log.Fatal("failed to get ui fs", zap.Error(err))
		}
	}

	// Start server
	wg.Add(1)
	defer wg.Done()

	go func() {
		<-ctx.Done()
		// When the app is shutting down, shutdown the HTTP server as well
		_ = srv.ec.Shutdown(context.Background())
	}()

	addr := fmt.Sprintf(":%d", cfg.API.PORT)

	if err := srv.ec.Start(addr); err != nil && err != http.ErrServerClosed {
		log.Error("http", zap.Error(err))
	} else {
		log.Info("Shutting down", zap.Error(err))
	}
}
