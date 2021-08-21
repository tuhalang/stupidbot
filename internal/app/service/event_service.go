package service

import (
	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

type EventService struct {
	config util.Config
	store  *db.Store
}

func NewEventService(config util.Config, store *db.Store) *EventService {

	service := &EventService{
		config: config,
		store:  store,
	}

	return service
}
