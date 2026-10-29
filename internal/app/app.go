package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/asfqx/SubTrack/config"
	"github.com/asfqx/SubTrack/pkg/httpserver"
	"github.com/asfqx/SubTrack/pkg/router"
	"github.com/rs/zerolog/log"
)

func Run(ctx context.Context, c config.Config) error {
	r := router.New()
	httpServer := httpserver.New(r, c.HttpServer)

	log.Info().Msg("App started!")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Info().Msg("App got signal to stop")

	httpServer.Close()

	return nil
}
