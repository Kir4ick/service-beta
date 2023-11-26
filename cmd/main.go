package main

import (
	"beta/internal/handler"
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
		Port: envReader.Get("PORT"),
	}

	databaseConfig := config.DatabaseConfig{
		Host: envReader.Get("DB_HOST"),
		Port: envReader.Get("DB_PORT"),
	}

	conf := config.NewConfig(&serverConfig, &databaseConfig)
	srv := new(server.Server)

	ctx := context.Background()
	databaseConnect := database.Connection(&databaseConfig, &ctx)

	handlers := new(handler.Handler)
	// вынес в анонимную функцию, так как нельзя написать if сразу после go
	go func() {
		if err := srv.Run(conf.Server.Port, handlers.InitRoutes()); err == nil {
			log.Fatalf("server error %s", err.Error())
		}
	}()

	shutdown(&ctx, srv, databaseConnect)
}
