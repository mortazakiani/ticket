package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mortazakiani/ticket/internal/api"
	db "github.com/mortazakiani/ticket/internal/db/sqlc"
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
	databseConectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
	config.DBUSERNAME,
	config.DBPASSWORD,
	config.DBHOST,
	config.DBPORT,
	config.DBDATABASE,
	)
	dbpool, err := pgxpool.New(context.Background(), databseConectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect  to  DB")
	}
	defer dbpool.Close()

	store := db.NewStore(dbpool)
	server ,err:= api.Newserver(config,store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect  to server ")
	}
}