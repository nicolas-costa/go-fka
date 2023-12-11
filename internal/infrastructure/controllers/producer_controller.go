package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gofka/internal/application"
	"net/http"
)

type ProducerController struct {
	messagingService application.Messenger
}

func NewProducerController(service application.Messenger) *ProducerController {
	return &ProducerController{
		messagingService: service,
	}
}

func (p *ProducerController) Ping(c *fiber.Ctx) error {
	p.messagingService.Post("Ping")

	return c.SendStatus(http.StatusNoContent)
}
