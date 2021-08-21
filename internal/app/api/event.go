package api

import (
	"log"
	"net/http"

	"github.com/tuhalang/stupidbot/internal/app/domain"

	"github.com/gin-gonic/gin"
)

func (server *Server) verifyToken(ctx *gin.Context) {
	token := ctx.Request.URL.Query().Get("hub.verify_token")
	challenge := ctx.Request.URL.Query().Get("hub.challenge")
	resp, err := server.eventService.VerifyToken(token, challenge)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.Data(http.StatusAccepted, "text", []byte(resp))
}

func (server *Server) handleEvent(ctx *gin.Context) {
	var req domain.InputMessage
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal("Server error: ", err)
		return
	}

	if err := server.eventService.HandleMessage(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusAccepted, nil)
}
