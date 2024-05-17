package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bahattincinic/fitwave/config"
	"github.com/bahattincinic/fitwave/database"
	"github.com/bahattincinic/fitwave/importer"
	"github.com/bahattincinic/fitwave/strava"
	pkgerrors "github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func main() {
	// More setup
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// Handle signals, cancel context
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGPIPE)
		for {
			sig, ok := <-ch
			if ok {
				if sig == syscall.SIGPIPE {
					log.Info("Got signal, ignoring...", zap.String("signal", sig.String()))
					continue
				}

				log.Info("Got signal, cleaning up...", zap.String("signal", sig.String()))
			} else {
				log.Info("Signal channel closed, cleaning up...")
			}
			cancelFunc()
			return
		}
	}()

	c, err := config.Parse()
	if err != nil {
		log.Fatal("init config failed", zap.Error(err))
		os.Exit(1)
	}

	{
		zlog, err := getLogger(c.Log)
		if err != nil {
			log.Fatal("init log failed", zap.Error(err))
			os.Exit(1)
		}
		log = zlog
	}

	// WorkGroup to coordinate start and shutdown all long running services
	wg := &sync.WaitGroup{}
	wg.Add(1) // for main

	db, err := database.NewDatabase(ctx, log.With(zap.String("module", "database")), c)
	if err != nil {
		log.Fatal("init database failed", zap.Error(err))
		os.Exit(1)
	}

	st := strava.NewStrava(ctx, c, log.With(zap.String("module", "strava")))
	im := importer.NewImporter(ctx, c, log.With(zap.String("module", "importer")), st, db)

	go func() {
		// Wait for all tasks to finish
		<-ctx.Done()
		wg.Done() // for main
	}()

	if err := im.Import(); err != nil {
		log.Fatal("import failed", zap.Error(err))
	}
}

func getLogger(log config.LogConfig) (*zap.Logger, error) {
	l, err := zapcore.ParseLevel(log.Level)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "ParseLevel")
	}

	var zc zap.Config
	if !config.Local() {
		zc = zap.NewProductionConfig()
		zc.Level = zap.NewAtomicLevelAt(l)
	} else {
		zc = zap.NewDevelopmentConfig()
		zc.Level = zap.NewAtomicLevelAt(l)
		zc.OutputPaths = []string{log.Output}
	}
	return zc.Build()
}
