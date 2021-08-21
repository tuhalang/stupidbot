package handler

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/domain"
)

func ProcessGuideAction(store *db.Store, script db.Script, senderId string) ([]byte, error) {
	replies, err := store.GetByScriptCode(context.Background(), script.Code)
	if err != nil {
		log.Print("Server error: ", err)
	}

	var quickReplies []domain.QuickReply
	for _, e := range replies {
		quickReplies = append(quickReplies, domain.QuickReply{
			ContentType: e.ContentType.String,
			Title:       e.Title.String,
			Payload:     e.Payload.String,
			ImageUrl:    e.ImageUrl.String,
		})
	}

	response := domain.ResponseMessage{
		Recipient: domain.Recipient{
			ID: senderId,
		},
		MessagingType: "RESPONSE",
		Message: domain.Message{
			Text:         script.MessageText.String,
			QuickReplies: quickReplies,
		},
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Marshal error: %s", err)
		return nil, err
	}

	return data, nil
}
