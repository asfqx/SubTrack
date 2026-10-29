package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/asfqx/SubTrack/config"
	"github.com/asfqx/SubTrack/internal/app"
	"github.com/asfqx/SubTrack/pkg/logger"
)

func main() {

	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	logger.Init(c.Logger)

	ctx := context.Background()

	err = app.Run(ctx, c)
	if err != nil {
		log.Error().Err(err).Msg("app.Run")
	}
}
