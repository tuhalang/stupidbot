package service

import (
	"context"

	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/service/handler"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func (eventService *EventService) handleProcess(senderId, text, postback, messageType string, script db.Script, fn func([]byte) error) error {

	var datas [][]byte
	var err error

	if util.MessagePostback == messageType {
		switch script.ActionProcessor {
		case util.ProcessPracticeAction:
			datas, _ = handler.ProcessPracticeAction(eventService.store, script, senderId, text, postback)
		case util.ProcessGuideAction:
			datas, _ = handler.ProcessGuideAction(eventService.store, script, senderId)
		case util.ProcessIntroAction:
			datas, _ = handler.ProcessIntroAction(eventService.store, script, senderId)
		case util.ProcessContributeAction:
			datas, _ = handler.ProcessContributeAction(eventService.store, script, senderId)
		case util.ProcessDonateAction:
			datas, _ = handler.ProcessDonateAction(eventService.store, script, senderId)
		case util.ProcessNextQuestionAction:
			datas, _ = handler.ProcessNextQuestionAction(eventService.store, script, senderId, text, postback)
		case util.ProcessRegisterMentorAction:
			datas, _ = handler.ProcessRegisterMentorAction(eventService.store, script, senderId)
		case util.ProcessAskMentorAction:
			datas, _ = handler.ProcessAskMentorAction(eventService.store, script, senderId, text, postback)
		}
	} else if util.MessageText == messageType {
		switch script.ActionProcessor {
		case util.ProcessGuideAction:
			datas, _ = handler.ProcessGuideAction(eventService.store, script, senderId)
		}
	}

	eventService.store.CreateSessionContext(context.Background(), db.CreateSessionContextParams{
		UserID:     senderId,
		ScriptCode: script.Code,
	})

	for _, data := range datas {
		err = fn(data)
	}

	return err
}
