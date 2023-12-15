// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gofka/internal/application"
	"gofka/internal/infrastructure/controllers"
	"gofka/internal/infrastructure/repositories"
	"gofka/internal/infrastructure/servers"
)

// Injectors from wire.go:

func initialize() *servers.FiberServer {
	kafkaRepository := repositories.NewKafkaRepository()
	healthService := application.NewHealthService(kafkaRepository)
	healthController := controllers.NewHealthController(healthService)
	pongService := application.NewPongService(kafkaRepository)
	consumerController := controllers.NewConsumerController(pongService)
	router := NewRouter(healthController, consumerController)
	fiberServer := servers.NewFiberServer(router)
	return fiberServer
}
