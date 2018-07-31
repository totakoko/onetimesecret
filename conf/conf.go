package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Configuration struct {
	ListenPort int    `default:"5000"`
	LogLevel   string `default:"info"`
	Store      StoreConfig
}

type StoreConfig struct {
	Addr      string `default:"127.0.0.1:6379"`
	Password  string `default:"password123"`
	Database  int    `default:"0"`
	Flush     bool   `default:"false"`
	KeyLength int    `default:"8"`
}

func New() Configuration {
	zerolog.TimeFieldFormat = ""
	var conf Configuration
	err := envconfig.Process("OTS", &conf)
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Msgf("Conf %+v", conf)

	return conf
}
