package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tuhalang/stupidbot/internal/app/domain"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func (eventService *EventService) HandleMessage(inputMessage domain.InputMessage) error {

	log.Print("EventMessage", inputMessage)

	for _, entry := range inputMessage.Entry {
		if len(entry.Messaging) == 0 {
			log.Print("No messages")
			return errors.New("no messages")
		}

		event := entry.Messaging[0]

		if len(event.Postback.Payload) > 0 {
			err := eventService.handlePostback(event.Sender.ID, event.Postback.Title, event.Postback.Payload)
			if err != nil {
				log.Print(err)
				return err
			}
		} else if len(event.Message.QuickReply.Payload) > 0 {
			err := eventService.handlePostback(event.Sender.ID, event.Message.Text, event.Message.QuickReply.Payload)
			if err != nil {
				log.Print(err)
				return err
			}
		} else if len(event.Message.Text) > 0 {
			if err := eventService.handleText(event.Sender.ID, event.Message.Text); err != nil {
				log.Print(err)
				return err
			}
		} else if len(event.Message.Attachments) > 0 {
			if err := eventService.handleMedia(event.Sender.ID, event.Message.Attachments); err != nil {
				log.Print(err)
				return err
			}
		}

	}

	return nil
}

func (eventService *EventService) handleMedia(senderId string, attachments []domain.ContentAttachment) error {
	conversation, err := eventService.store.GetConversationByUser(context.Background(), senderId)
	if err == nil {
		for _, attachment := range attachments {

			response := domain.ResponseAttachment{
				Recipient: domain.Recipient{
					ID: conversation.MentorID,
				},
				Message: domain.Attachment{
					Attachment: attachment,
				},
			}

			data, err := json.Marshal(response)
			if err != nil {
				log.Printf("Marshal error: %s", err)
				return err
			}
			eventService.sendRequest(data)
		}
		return nil
	}

	conversation, err = eventService.store.GetConversationByMentor(context.Background(), senderId)
	if err == nil {
		for _, attachment := range attachments {
			response := domain.ResponseAttachment{
				Recipient: domain.Recipient{
					ID: conversation.UserID,
				},
				Message: domain.Attachment{
					Attachment: attachment,
				},
			}

			data, err := json.Marshal(response)
			if err != nil {
				log.Printf("Marshal error: %s", err)
				return err
			}
			eventService.sendRequest(data)
		}
		return nil
	}

	script, _ := eventService.store.GetScriptByCode(context.Background(), util.DefaultScript)
	return eventService.handleProcess(senderId, "", "", util.MessageText, script, eventService.sendRequest)
}

// Dinh dang tin nhan postback: COMMAND|param...
func (eventService *EventService) handlePostback(senderId, title, postback string) error {
	log.Print("message postback from client: [title=" + title + ", postback=" + postback + "]")
	commands := strings.Split(postback, "|")
	script, err := eventService.store.GetScriptByCode(context.Background(), strings.ToUpper(commands[0]))

	if err != nil {
		log.Printf("Database error %v", err)
		script, _ = eventService.store.GetScriptByCode(context.Background(), util.DefaultScript)
	}

	return eventService.handleProcess(senderId, title, postback, util.MessagePostback, script, eventService.sendRequest)
}

func (eventService *EventService) handleText(senderId string, text string) error {
	log.Print("message text from client: " + text)

	conversation, err := eventService.store.GetConversationByUser(context.Background(), senderId)
	if err == nil {
		if text == util.ExitComman {
			err := eventService.store.InvalidConversation(context.Background(), conversation.ID)
			if err != nil {
				log.Print(err)
				return err
			}
		} else {
			response := domain.ResponseMessage{
				Recipient: domain.Recipient{
					ID: conversation.MentorID,
				},
				MessagingType: "RESPONSE",
				Message: domain.Message{
					Text: text,
				},
			}

			data, err := json.Marshal(response)
			if err != nil {
				log.Printf("Marshal error: %s", err)
				return err
			}
			eventService.sendRequest(data)
			return nil
		}
	}

	conversation, err = eventService.store.GetConversationByMentor(context.Background(), senderId)
	if err == nil {
		if text == util.ExitComman {
			err := eventService.store.InvalidConversation(context.Background(), conversation.ID)
			if err != nil {
				log.Print(err)
				return err
			}
		} else {
			response := domain.ResponseMessage{
				Recipient: domain.Recipient{
					ID: conversation.UserID,
				},
				MessagingType: "RESPONSE",
				Message: domain.Message{
					Text: text,
				},
			}

			data, err := json.Marshal(response)
			if err != nil {
				log.Printf("Marshal error: %s", err)
				return err
			}
			eventService.sendRequest(data)
			return nil
		}
	}

	script, err := eventService.store.GetScriptByCode(context.Background(), strings.ToUpper(text))
	if err != nil {
		log.Printf("Database error %v", err)
		script, _ = eventService.store.GetScriptByCode(context.Background(), util.DefaultScript)
	}

	return eventService.handleProcess(senderId, text, "", util.MessageText, script, eventService.sendRequest)
}

func (eventService *EventService) sendRequest(data []byte) error {
	uri := fmt.Sprintf("%s%s", eventService.config.FacebookApi, eventService.config.PageAccessToken)
	req, err := http.NewRequest(
		"POST",
		uri,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		log.Printf("MESSAGE: %#v", res)
	}

	return nil
}
