package helpers

import (
	"github.com/rs/zerolog/log"
)

func Try(err error) {
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
