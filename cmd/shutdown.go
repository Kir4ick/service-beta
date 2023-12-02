package main

import (
	"beta/pkg/database"
	"beta/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func shutdown(ctx *context.Context, srv *server.Server, db *database.DatabaseClient, cancel context.CancelFunc) {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	defer cancel()

	log.Println("shutdown connect to database")
	if err := db.Disconnect(ctx); err != nil {
		log.Fatalf("error to disconnect to database %s", err.Error())
	}

	log.Println("shutdown server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error to shutdown server %s", err.Error())
	}

}
