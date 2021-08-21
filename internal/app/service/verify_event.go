package service

import (
	"errors"
	"log"
)

func (eventService *EventService) VerifyToken(token string, challenge string) (string, error) {
	log.Print("verify token: [token: " + token + ", challange: " + challenge + "]")

	if eventService.config.VerifyToken == token {
		return challenge, nil
	}
	return "", errors.New("token invalid")
}
