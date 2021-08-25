package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	db "github.com/tuhalang/stupidbot/internal/app/db/sqlc"
	"github.com/tuhalang/stupidbot/internal/app/domain"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func ProcessRegisterMentorAction(store *db.Store, script db.Script, senderId string) ([][]byte, error) {
	var resp [][]byte

	_, err := store.CreateUser(context.Background(), db.CreateUserParams{
		ID: senderId,
		IsMentor: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Status: 0,
	})
	if err != nil {
		log.Print("CreateUser error: ", err)
	}

	replies, err := store.GetByScriptCode(context.Background(), util.DefaultScript)
	if err != nil {
		log.Print("GetByScriptCode error: ", err)
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

func ProcessAskMentorAction(store *db.Store, script db.Script, senderId, text, postback string) ([][]byte, error) {
	log.Print("Start: ProcessAskMentorAction")
	// ASK_MENTOR|PRACTICE|VAT_LY|VAT_LY_CH1_DDC|1|A
	var resp [][]byte

	mentor, err := store.GetMentorAvailable(context.Background())
	if err != nil {
		log.Print(err)
		responseUser := domain.ResponseMessage{
			Recipient: domain.Recipient{
				ID: senderId,
			},
			MessagingType: "RESPONSE",
			Message: domain.Message{
				Text: "Xin lỗi hiện tại không có mentor nào đang online.",
			},
		}

		dataUser, err := json.Marshal(responseUser)
		if err != nil {
			log.Printf("Marshal error: %s", err)
			return nil, err
		}
		resp = append(resp, dataUser)

		return resp, nil
	}

	_, err = store.CreateConversation(context.Background(), db.CreateConversationParams{
		UserID:   senderId,
		MentorID: mentor.ID,
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// ================================= //
	commands := strings.Split(postback, "|")
	questionId, err := strconv.ParseInt(commands[4], 10, 64)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	question, err := store.GetQuestionById(context.Background(), questionId)
	if err != nil {
		log.Print("GetQuestionById error: ", err)
		return nil, err
	}
	responseImageQuestion := domain.ResponseAttachment{
		Recipient: domain.Recipient{
			ID: mentor.ID,
		},
		Message: domain.Attachment{},
	}

	responseImageQuestion.Message.Attachment.Type = "image"
	responseImageQuestion.Message.Attachment.Payload.URL = question.ContentImage
	dataQuestion, err := json.Marshal(responseImageQuestion)
	if err != nil {
		log.Printf("Marshal error: %s", err)
		return nil, err
	}
	resp = append(resp, dataQuestion)

	responseMentor := domain.ResponseMessage{
		Recipient: domain.Recipient{
			ID: mentor.ID,
		},
		MessagingType: "RESPONSE",
		Message: domain.Message{
			Text: "Chào bạn, bạn vui lòng hướng dẫn mình bài tập này.",
		},
	}

	dataMentor, err := json.Marshal(responseMentor)
	if err != nil {
		log.Printf("Marshal error: %s", err)
		return nil, err
	}
	resp = append(resp, dataMentor)

	// ===================================== //

	responseUser := domain.ResponseMessage{
		Recipient: domain.Recipient{
			ID: senderId,
		},
		MessagingType: "RESPONSE",
		Message: domain.Message{
			Text: "Bạn đã được kết nối tói mentor, vui lòng đợi phản hồi. Bạn có thể nhập [q] để thoát khỏi chế độ này.",
		},
	}

	dataUser, err := json.Marshal(responseUser)
	if err != nil {
		log.Printf("Marshal error: %s", err)
		return nil, err
	}
	resp = append(resp, dataUser)

	return resp, nil
}
