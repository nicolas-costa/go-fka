//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gofka/internal/application"
	"gofka/internal/infrastructure/controllers"
	"gofka/internal/infrastructure/servers"
)

func initialize() *servers.FiberServer {
	wire.Build(
		servers.NewFiberServer,
		controllers.NewHealthController,
		NewRouter,
		wire.Bind(new(servers.Router), new(*Router)),
		application.NewHealthService,
		wire.Bind(new(application.Checker), new(*application.HealthService)),
	)

	return &servers.FiberServer{}
}
