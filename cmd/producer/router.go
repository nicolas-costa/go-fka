package main

import (
	"gofka/internal/infrastructure/controllers"
	"gofka/internal/infrastructure/servers"
)

type Router struct {
	healthController   *controllers.HealthController
	producerController *controllers.ProducerController
}

func NewRouter(healthController *controllers.HealthController,
	producerController *controllers.ProducerController) *Router {
	return &Router{
		healthController:   healthController,
		producerController: producerController,
	}
}

func (r *Router) SetRoutes(app *servers.FiberServer) {
	app.Get("/health", r.healthController.Check)

	app.Get("ping", r.producerController.Ping)
}
