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

func ProcessPracticeAction(store *db.Store, script db.Script, senderId, text, postback string) ([][]byte, error) {
	// PRACTICE|VAT_LY|VAT_LY_CH1_DDC

	commands := strings.Split(postback, "|")
	var resp [][]byte

	// PRACTICE
	if len(commands) == 1 {
		subjects, err := store.GetSubjects(context.Background())
		if err != nil {
			log.Print("ListTopics error: ", err)
			return nil, err
		}

		var quickReplies []domain.QuickReply
		for _, e := range subjects {
			quickReplies = append(quickReplies, domain.QuickReply{
				ContentType: "text",
				Title:       e.SubjectName,
				Payload:     postback + "|" + e.SubjectCode,
				ImageUrl:    e.Image.String,
			})
		}

		response := domain.ResponseMessage{
			Recipient: domain.Recipient{
				ID: senderId,
			},
			MessagingType: "RESPONSE",
			Message: domain.Message{
				Text:         "Chọn môn học",
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

		// PRACTICE|VAT_LY
	} else if len(commands) == 2 {
		topics, err := store.ListTopics(context.Background(), commands[1])
		if err != nil {
			log.Print("ListTopics error: ", err)
			return nil, err
		}

		var quickReplies []domain.QuickReply
		for _, e := range topics {
			quickReplies = append(quickReplies, domain.QuickReply{
				ContentType: "text",
				Title:       e.TopicName.String,
				Payload:     postback + "|" + e.TopicCode,
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
		resp = append(resp, data)
		return resp, nil

		// PRACTICE|VAT_LY|VAT_LY_CH1_DDC
	} else if len(commands) == 3 {

		question, err := store.GetRandomQuestion(context.Background(), db.GetRandomQuestionParams{
			TopicCode:   commands[2],
			SubjectCode: commands[1],
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
				Payload:     postback + "|" + strconv.Itoa(int(question.ID)) + "|" + e.Payload,
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

		// PRACTICE|VAT_LY|VAT_LY_CH1_DDC|1|A
	} else if len(commands) == 5 {
		questionId, err := strconv.ParseInt(commands[3], 10, 64)
		if err != nil {
			log.Print(err)
			return nil, err
		}

		question, err := store.GetQuestionById(context.Background(), questionId)
		if err != nil {
			log.Print("GetQuestionById error: ", err)
			return nil, err
		}

		response := domain.ResponseMessage{
			Recipient: domain.Recipient{
				ID: senderId,
			},
			MessagingType: "RESPONSE",
			Message:       domain.Message{},
		}

		if question.CorrectAnswer == commands[4] {
			response.Message.Text = "CHÍNH XÁC đáp án đúng là: " + question.CorrectAnswer
		} else {
			response.Message.Text = "CHƯA CHÍNH XÁC đáp án đúng là: " + question.CorrectAnswer
		}

		options, err := store.GetByScriptCode(context.Background(), util.ExtraOptionsScript)
		if err != nil {
			log.Print("Server error: ", err)
			return nil, err
		}

		var quickReplies []domain.QuickReply
		for _, e := range options {
			quickReplies = append(quickReplies, domain.QuickReply{
				ContentType: e.ContentType,
				Title:       e.Title,
				Payload:     e.Payload + "|" + postback,
			})
		}
		response.Message.QuickReplies = quickReplies

		data, err := json.Marshal(response)
		if err != nil {
			log.Printf("Marshal error: %s", err)
			return nil, err
		}
		resp = append(resp, data)

		return resp, nil
	}
	return nil, nil
}
