package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123}).With().Timestamp().Logger()
	logger.Error().Stack().Msg("An error occured")
	logger.Trace().Stack().Msg("exception")
	logger.Info().Int("asd", 2).Bool("sd", false).Stack()
	// logger.Panic().Msg("error")
}
