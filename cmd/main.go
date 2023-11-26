package main

import (
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

	// вынес в анонимную функцию, так как нельзя написать if сразу после go
	go func(conf *config.Config) {
		if err := srv.Run(conf.Server.Port); err == nil {
			log.Fatalf("server error %s", err.Error())
		}
	}(conf)

	shutdown(&ctx, srv, databaseConnect)
}
