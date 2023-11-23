package main

import (
	"beta/pkg/config"
	"beta/pkg/reader"
	"beta/server"
	"log"
)

func main() {
	envReader := new(reader.Env)
	envReader = envReader.NewEnv()

	serverConfig := config.ServerConfig{
		Port: envReader.Get("PORT"),
	}

	databaseConfig := config.DatabaseConfig{
		Connection:   "mongodb",
		DatabaseName: envReader.Get("DB_NAME"),
		Host:         envReader.Get("DB_HOST"),
		Port:         envReader.Get("DB_PORT"),
		User:         envReader.Get("DB_USER"),
		Password:     envReader.Get("DB_PASSWORD"),
	}

	conf := config.NewConfig(&serverConfig, &databaseConfig)
	srv := new(server.Server)

	if err := srv.Run(conf.Server.Port); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
