package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gofka/internal/application"
	"net/http"
)

type ProducerController struct {
	pingerService application.Pinger
}

func NewProducerController(service application.Pinger) *ProducerController {
	return &ProducerController{
		pingerService: service,
	}
}

func (p *ProducerController) Ping(c *fiber.Ctx) error {
	p.pingerService.Ping()

	return c.SendStatus(http.StatusNoContent)
}
