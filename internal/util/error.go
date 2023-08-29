package util

import (
	"github.com/rs/zerolog/log"
)

func FatalIfError(err error) {
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
