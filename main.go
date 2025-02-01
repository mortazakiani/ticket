package main

import (
	"os"

	"github.com/mortazakiani/ticket/internal/utiles"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//load configs
	config,err := utiles.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config,lunch app faild")
	}

	if config.APPDEBUG == "true" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

}