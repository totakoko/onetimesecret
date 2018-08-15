package main

import (
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.dreau.fr/home/onetimesecret/conf"
	"gitlab.dreau.fr/home/onetimesecret/httpserver"
	"gitlab.dreau.fr/home/onetimesecret/store"
)

func main() {
	err := startServer()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func startServer() error {
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
	return server.Run(":" + strconv.Itoa(config.ListenPort))
}
