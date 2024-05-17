package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bahattincinic/fitwave/config"
	"github.com/bahattincinic/fitwave/database"
	"github.com/bahattincinic/fitwave/importer"
	"github.com/bahattincinic/fitwave/strava"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func main() {
	var (
		flagConfigPath string
	)
	{
		flag.StringVar(&flagConfigPath, "config-path", "", "Configuration path")
	}
	flagLogLevel := zap.LevelFlag("l", zapcore.DebugLevel, "Log level")
	flag.Parse()

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

	c, err := config.Parse(flagConfigPath)
	if err != nil {
		log.Fatal("init config failed", zap.Error(err))
		os.Exit(1)
	}

	if !config.Local() {
		{
			c := zap.NewProductionConfig()
			c.Level = zap.NewAtomicLevelAt(*flagLogLevel)
			zlog, err := c.Build()
			if err != nil {
				panic(err)
			}
			log = zlog
		}
	} else {
		{
			zc := zap.NewDevelopmentConfig()
			zc.Level = zap.NewAtomicLevelAt(*flagLogLevel)
			zc.OutputPaths = []string{c.Log.Output}

			zlog, err := zc.Build()
			if err != nil {
				panic(err)
			}
			log = zlog
		}
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
