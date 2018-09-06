package main

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/totakoko/onetimesecret/conf"
	"gitlab.com/totakoko/onetimesecret/helpers"
	"gitlab.com/totakoko/onetimesecret/httpserver"
	"gitlab.com/totakoko/onetimesecret/store"
)

func main() {
	helpers.TryOrFatal(startServer())
}

func startServer() error {
	zerolog.TimeFieldFormat = ""
	log.Info().Msg("Starting server")

	config, err := conf.New()
	if err != nil {
		return err
	}

	logLevel, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(logLevel)

	store := store.New(config.Store)
	if err := store.Init(); err != nil {
		return err
	}

	server := &httpserver.HTTPServer{
		PublicURL: config.PublicURL,
		Store:     store,
	}
	if err := server.Init(); err != nil {
		return err
	}

	go func() {
		err := server.Run(":" + strconv.Itoa(config.ListenPort))
		if err != http.ErrServerClosed {
			log.Fatal().Msg(err.Error())
		}
		log.Warn().Msgf("Server stopped")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Warn().Msgf("Shutting down server...")
	if err := server.Shutdown(); err != nil {
		return err
	}

	log.Warn().Msgf("Server exiting")
	return nil
}
