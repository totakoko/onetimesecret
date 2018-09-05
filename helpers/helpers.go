package helpers

import (
	"github.com/rs/zerolog/log"
)

func LogOnError(err error) {
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
func TryOrFatal(err error) {
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
