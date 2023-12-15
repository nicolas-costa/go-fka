package main

import (
	"gofka/internal/infrastructure/controllers"
	"gofka/internal/infrastructure/servers"
)

type Router struct {
	healthController   *controllers.HealthController
	consumerController *controllers.ConsumerController
}

func NewRouter(healthController *controllers.HealthController,
	consumerController *controllers.ConsumerController) *Router {
	return &Router{
		healthController:   healthController,
		consumerController: consumerController,
	}
}

func (r *Router) SetRoutes(app *servers.FiberServer) {
	app.Get("/health", r.healthController.Check)

	messaging := app.Group("messaging")

	messaging.Get("/pong", r.consumerController.Pong)
}
