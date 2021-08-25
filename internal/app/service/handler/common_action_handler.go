package handler

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/domain"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func ProcessGuideAction(store *db.Store, script db.Script, senderId string) ([][]byte, error) {

	var resp [][]byte

	replies, err := store.GetByScriptCode(context.Background(), script.Code)
	if err != nil {
		log.Print("Server error: ", err)
	}

	var quickReplies []domain.QuickReply
	for _, e := range replies {
		quickReplies = append(quickReplies, domain.QuickReply{
			ContentType: e.ContentType,
			Title:       e.Title,
			Payload:     e.Payload,
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

	resp = append(resp, data)
	return resp, nil
}

func ProcessIntroAction(store *db.Store, script db.Script, senderId string) ([][]byte, error) {
	var resp [][]byte

	replies, err := store.GetByScriptCode(context.Background(), script.Code)
	if err != nil {
		log.Print("Server error: ", err)
	}

	var quickReplies []domain.QuickReply
	for _, e := range replies {
		quickReplies = append(quickReplies, domain.QuickReply{
			ContentType: e.ContentType,
			Title:       e.Title,
			Payload:     e.Payload,
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

	resp = append(resp, data)
	return resp, nil
}

func ProcessContributeAction(store *db.Store, script db.Script, senderId string) ([][]byte, error) {
	var resp [][]byte

	replies, err := store.GetByScriptCode(context.Background(), util.DefaultScript)
	if err != nil {
		log.Print("Server error: ", err)
	}

	var quickReplies []domain.QuickReply
	for _, e := range replies {
		quickReplies = append(quickReplies, domain.QuickReply{
			ContentType: e.ContentType,
			Title:       e.Title,
			Payload:     e.Payload,
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

	resp = append(resp, data)
	return resp, nil
}

func ProcessDonateAction(store *db.Store, script db.Script, senderId string) ([][]byte, error) {
	var resp [][]byte

	replies, err := store.GetByScriptCode(context.Background(), util.DefaultScript)
	if err != nil {
		log.Print("Server error: ", err)
	}

	var quickReplies []domain.QuickReply
	for _, e := range replies {
		quickReplies = append(quickReplies, domain.QuickReply{
			ContentType: e.ContentType,
			Title:       e.Title,
			Payload:     e.Payload,
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

	resp = append(resp, data)
	return resp, nil
}
