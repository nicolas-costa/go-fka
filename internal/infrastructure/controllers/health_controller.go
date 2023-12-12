package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gofka/internal/application"
)

type HealthController struct {
	healthService application.Checker
}

func NewHealthController(service application.Checker) *HealthController {
	return &HealthController{
		healthService: service,
	}
}

func (h *HealthController) Check(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"kafka": h.healthService.Check(),
	})
}
