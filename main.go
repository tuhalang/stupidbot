package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tuhalang/stupidbot/internal/app/api"
	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/service"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	eventService := service.NewEventService(config, store)

	server, err := api.NewServer(eventService)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
