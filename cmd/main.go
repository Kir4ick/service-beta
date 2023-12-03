package main

import (
	"beta/internal/handler"
	repository "beta/internal/repositories"
	service "beta/internal/services"
	"beta/pkg/config"
	"beta/pkg/database"
	"beta/pkg/reader"
	"beta/pkg/regulation"
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
		Host:     envReader.Get("DB_HOST"),
		Port:     envReader.Get("DB_PORT"),
		Name:     envReader.Get("DB_NAME"),
		User:     envReader.Get("DB_USER"),
		Password: envReader.Get("DB_PASSWORD"),
	}

	conf := config.NewConfig(&serverConfig, &databaseConfig).SetGammaUrl(
		envReader.Get("GAMMA_SERVICE_URL"))
	srv := new(server.Server)

	ctx, cancel := context.WithCancel(context.Background())

	requestRegulation := regulation.NewRequestRegulation()
	databaseConnect := database.Connection(&databaseConfig, &ctx)
	var repositories = repository.NewRepository(databaseConnect, databaseConfig)
	var voteRepository repository.IVotingRepository = repository.NewVotingRepository(repositories)
	var services service.IService = service.NewService(voteRepository, conf, requestRegulation)
	handlers := handler.NewHandler(services, &ctx)

	go func() {
		if err := srv.Run(conf.Server.Port, handlers.InitRoutes()); err == nil {
			log.Fatalf("server error %s", err.Error())
		}
	}()

	services.ClearInfoForRequests(requestRegulation)

	shutdown(&ctx, srv, databaseConnect, cancel)
}
