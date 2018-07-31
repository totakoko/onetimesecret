package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.dreau.fr/home/onetimesecret/conf"
	"gitlab.dreau.fr/home/onetimesecret/httpserver"
	"gitlab.dreau.fr/home/onetimesecret/store"
)

func main() {
	log.Info().Msg("Starting server")

	rand.Seed(int64(time.Now().Nanosecond()))

	config := conf.New()
	logLevel, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal().Err(err)
	}
	zerolog.SetGlobalLevel(logLevel)
	store := store.New(config.Store)
	err = store.Init()
	if err != nil {
		log.Fatal().Err(err)
	}

	server := &httpserver.HTTPServer{Store: store}
	server.Init()
	err = server.Run(":" + strconv.Itoa(config.ListenPort))
	if err != nil {
		log.Fatal().Err(err)
	}
}
