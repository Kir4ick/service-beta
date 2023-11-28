package main

import (
	"beta/internal/handler"
	repository "beta/internal/repositories"
	service "beta/internal/services"
	"beta/pkg/config"
	"beta/pkg/database"
	"beta/pkg/reader"
	"beta/server"
	"context"
	"log"
)

func main() {
	envReader := reader.GetEnvReader()

	serverConfig := config.ServerConfig{
		Port: envReader.Get("HTTP_PORT"),
	}

	databaseConfig := config.DatabaseConfig{
		Host: envReader.Get("DB_HOST"),
		Port: envReader.Get("DB_PORT"),
	}

	conf := config.NewConfig(&serverConfig, &databaseConfig).SetGammaUrl(
		envReader.Get("GAMMA_SERVICE_URL"))
	srv := new(server.Server)

	ctx := context.Background()
	databaseConnect := database.Connection(&databaseConfig, &ctx)
	repositories := repository.NewRepository(databaseConnect)
	services := service.NewService(repositories, conf, &ctx)
	handlers := handler.NewHandler(services)

	go func() {
		if err := srv.Run(conf.Server.Port, handlers.InitRoutes()); err == nil {
			log.Fatalf("server error %s", err.Error())
		}
	}()

	shutdown(&ctx, srv, databaseConnect)
}
