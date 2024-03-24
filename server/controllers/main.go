package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/damocles217/microservice/server/broker"
	"github.com/damocles217/microservice/server/services"
	"github.com/gin-gonic/gin"
)

type Send struct {
	Message string `json:"message"`
}

type Handler struct {
	Services *services.Services
}

func SetNewHandler(services *services.Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) GetUser(ctx *gin.Context) {
	message := Send{
		Message: "Hola mundo desde kafka",
	}

	jsonBytes, err := json.Marshal(message)

	if err != nil {
		fmt.Println(err)
	}

	broker.PushMessageToQueue("user_video", jsonBytes, 0)

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": strconv.Itoa(h.Services.GetUser()),
	})
}
