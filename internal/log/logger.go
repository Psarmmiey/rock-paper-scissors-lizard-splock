package log

import "github.com/rs/zerolog/log"

func LogInfo(message interface{}) {
	log.Info().Msgf("%v ", message)
}

func LogDebug(message interface{}) {
	log.Debug().Msgf("%v ", message)
}
