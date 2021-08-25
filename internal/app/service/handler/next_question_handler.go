package handler

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/domain"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func ProcessNextQuestionAction(store *db.Store, script db.Script, senderId, text, postback string) ([][]byte, error) {
	log.Print("Start: ProcessNextQuestionAction")
	// NEXT_QUESTION|PRACTICE|VAT_LY|VAT_LY_CH1_DDC|1|A
	commands := strings.Split(postback, "|")
	var resp [][]byte

	if commands[1] == "PRACTICE" {
		question, err := store.GetRandomQuestion(context.Background(), db.GetRandomQuestionParams{
			TopicCode:   commands[3],
			SubjectCode: commands[2],
		})
		if err != nil {
			log.Print("GetRandomQuestion error: ", err)
			return nil, err
		}

		options, err := store.GetByScriptCode(context.Background(), util.OptionsScript)
		if err != nil {
			log.Print("GetByScriptCode error: ", err)
			return nil, err
		}

		var quickReplies []domain.QuickReply
		for _, e := range options {
			quickReplies = append(quickReplies, domain.QuickReply{
				ContentType: e.ContentType,
				Title:       e.Title,
				Payload:     commands[1] + "|" + commands[2] + "|" + commands[3] + "|" + strconv.Itoa(int(question.ID)) + "|" + e.Payload,
			})
		}

		response := domain.ResponseAttachment{
			Recipient: domain.Recipient{
				ID: senderId,
			},
			Message: domain.Attachment{},
		}

		response.Message.Attachment.Type = "image"
		response.Message.Attachment.Payload.URL = question.ContentImage
		data, err := json.Marshal(response)
		if err != nil {
			log.Printf("Marshal error: %s", err)
			return nil, err
		}
		resp = append(resp, data)

		responseOptions := domain.ResponseMessage{
			Recipient: domain.Recipient{
				ID: senderId,
			},
			MessagingType: "RESPONSE",
			Message: domain.Message{
				Text:         "[" + strconv.Itoa(int(question.ID)) + "] " + question.ContentText,
				QuickReplies: quickReplies,
			},
		}
		dataOptions, err := json.Marshal(responseOptions)
		if err != nil {
			log.Printf("Marshal error: %s", err)
			return nil, err
		}
		resp = append(resp, dataOptions)

		return resp, nil
	} else if commands[1] == "TEST" {

	}

	return nil, nil
}
