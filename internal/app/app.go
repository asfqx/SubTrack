package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/asfqx/SubTrack/config"
	"github.com/asfqx/SubTrack/pkg/httpserver"
	"github.com/asfqx/SubTrack/pkg/pgPool"
	"github.com/asfqx/SubTrack/pkg/router"

	"github.com/rs/zerolog/log"
)

func Run(ctx context.Context, c config.Config) error {
	postgresAdapter, err := pgPool.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgresPool.New: %w", err)
	}

	r := router.New()
	httpServer := httpserver.New(r, c.HttpServer)

	log.Info().Msg("App started!")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Info().Msg("App got signal to stop")

	httpServer.Close()
	postgresAdapter.Close()

	log.Info().Msg("App stopped!")

	return nil
}
