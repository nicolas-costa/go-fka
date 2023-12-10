package main

import (
	"gofka/internal/infrastructure/controllers"
	"gofka/internal/infrastructure/servers"
)

type Router struct {
	healthController *controllers.HealthController
}

func NewRouter(healthController *controllers.HealthController) *Router {
	return &Router{
		healthController: healthController,
	}
}

func (r *Router) SetRoutes(app *servers.FiberServer) {
	app.Get("/health", r.healthController.Check)
}
