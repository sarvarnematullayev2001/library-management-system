package main

import (
	"hw/prac/library_api_gateway/api"
	"hw/prac/library_api_gateway/config"
	"hw/prac/library_api_gateway/pkg/logger"
	"hw/prac/library_api_gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "library_api_gateway")

	gprcClients, _ := services.NewGrpcClients(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	server.Run(cfg.HttpPort)
}
