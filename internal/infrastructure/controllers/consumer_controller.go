package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gofka/internal/application"
)

type ConsumerController struct {
	pongerService application.Ponger
}

func NewConsumerController(service application.Ponger) *ConsumerController {
	return &ConsumerController{
		pongerService: service,
	}
}

func (p *ConsumerController) Pong(c *fiber.Ctx) error {
	messages := p.pongerService.Pong()

	return c.JSON(map[string]interface{}{
		"messages": messages,
	})
}
