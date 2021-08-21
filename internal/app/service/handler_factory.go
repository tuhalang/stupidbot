package service

import (
	"context"

	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/service/handler"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func (eventService *EventService) handleProcess(senderId, text, postback, messageType string, script db.Script, fn func([]byte) error) error {

	var data []byte

	if util.MessagePostback == messageType {
		switch script.ActionProcessor {
		case util.ProcessPracticeAction:
			data, _ = handler.ProcessPracticeAction(eventService.store, script, senderId, text, postback)
		}
	} else if util.MessageText == messageType {
		switch script.ActionProcessor {
		case util.ProcessGuideAction:
			data, _ = handler.ProcessGuideAction(eventService.store, script, senderId)
		}
	}

	eventService.store.CreateSessionContext(context.Background(), db.CreateSessionContextParams{
		UserID:     senderId,
		ScriptCode: script.Code,
	})

	return fn(data)
}
