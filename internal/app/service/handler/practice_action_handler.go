package handler

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/domain"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func ProcessPracticeAction(store *db.Store, script db.Script, senderId, text, postback string) ([]byte, error) {
	// PRACTICE|VATLY|VAT_LY_CH1_DDDH
	// PRACTICE
	commands := strings.Split(postback, "|")
	if len(commands) == 1 {
		topics, err := store.ListTopics(context.Background(), util.DefaultSubject)
		if err != nil {
			log.Print("Server error: ", err)
			return nil, err
		}

		var quickReplies []domain.QuickReply
		for _, e := range topics {
			quickReplies = append(quickReplies, domain.QuickReply{
				ContentType: "text",
				Title:       e.TopicName.String,
				Payload:     commands[0] + "|" + util.DefaultSubject + "|" + e.TopicCode,
				ImageUrl:    "https://drive.google.com/file/d/1rx4iRQ9STOKEuqg9ERV1w262ikvjn5-8/view?usp=sharing",
			})
		}

		response := domain.ResponseMessage{
			Recipient: domain.Recipient{
				ID: senderId,
			},
			MessagingType: "RESPONSE",
			Message: domain.Message{
				Text:         "Chọn chủ đề để luyện tập",
				QuickReplies: quickReplies,
			},
		}

		data, err := json.Marshal(response)
		if err != nil {
			log.Printf("Marshal error: %s", err)
			return nil, err
		}
		return data, nil

	} else if len(commands) == 2 {
		topics, err := store.ListTopics(context.Background(), commands[1])
		if err != nil {
			log.Print("Server error: ", err)
			return nil, err
		}

		var quickReplies []domain.QuickReply
		for _, e := range topics {
			quickReplies = append(quickReplies, domain.QuickReply{
				ContentType: "text",
				Title:       e.TopicName.String,
				Payload:     commands[0] + "|" + commands[1] + "|" + e.TopicCode,
				ImageUrl:    e.Image.String,
			})
		}

		response := domain.ResponseMessage{
			Recipient: domain.Recipient{
				ID: senderId,
			},
			MessagingType: "RESPONSE",
			Message: domain.Message{
				Text:         "Chọn chủ đề để luyện tập",
				QuickReplies: quickReplies,
			},
		}

		data, err := json.Marshal(response)
		if err != nil {
			log.Printf("Marshal error: %s", err)
			return nil, err
		}
		return data, nil
	} else if len(commands) == 3 {
		topics, err := store.ListTopics(context.Background(), commands[1])
		if err != nil {
			log.Print("Server error: ", err)
			return nil, err
		}

		var quickReplies []domain.QuickReply
		for _, e := range topics {
			quickReplies = append(quickReplies, domain.QuickReply{
				ContentType: "text",
				Title:       e.TopicName.String,
				Payload:     commands[0] + "|" + commands[1] + e.TopicCode,
			})
		}

		response := domain.ResponseMessage{
			Recipient: domain.Recipient{
				ID: senderId,
			},
			MessagingType: "RESPONSE",
			Message: domain.Message{
				Text:         "Chọn chủ đề để luyện tập",
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
	return nil, nil
}
