//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gofka/internal/application"
	"gofka/internal/infrastructure/controllers"
	"gofka/internal/infrastructure/repositories"
	"gofka/internal/infrastructure/servers"
)

func initialize() *servers.FiberServer {
	wire.Build(
		servers.NewFiberServer,
		controllers.NewHealthController,
		controllers.NewConsumerController,
		repositories.NewKafkaRepository,
		NewRouter,
		application.NewHealthService,
		application.NewPongService,
		wire.Bind(new(servers.Router), new(*Router)),
		wire.Bind(new(application.Checker), new(*application.HealthService)),
		wire.Bind(new(application.Healther), new(*repositories.KafkaRepository)),
		wire.Bind(new(application.Ponger), new(*application.PongService)),
		wire.Bind(new(application.Receiver), new(*repositories.KafkaRepository)),
	)

	return &servers.FiberServer{}
}
