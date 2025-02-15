package main

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mortazakiani/ticket/internal/api"
	db "github.com/mortazakiani/ticket/internal/db/sqlc"
	"github.com/mortazakiani/ticket/internal/routes"
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

	runDBMigration(config.MIGRATIONURL,databseConectionString)

	store := db.NewStore(dbpool)
	server ,err:= api.Newserver(config,store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect  to server ")
	}
	
	if err := routes.SetupRoutes(server); err != nil {
		log.Fatal().Err(err).Msg("cannot  set routes")
	} 

	err=server.Start(config.APPPORT)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server ")
	}
}


func runDBMigration(migrationURL string, dbsource string ){
	
	m, err := migrate.New(migrationURL,dbsource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create  new migration ")
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("cannot run up ")
	}


	log.Info().Msg("migration is done ")
}