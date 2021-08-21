package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tuhalang/stupidbot/internal/app/domain"
	"github.com/tuhalang/stupidbot/internal/app/util"
)

func (eventService *EventService) HandleMessage(inputMessage domain.InputMessage) error {

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
		}

	}

	return nil
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
