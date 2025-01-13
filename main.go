package main

import (
	"context"
	"go_server/server"
	"go_server/server/handlers"
	"go_server/storage"
	"log"
)

const (
	localhost         = ":80"
	storageSqlitePath = "data/sqlite/storage.db"
)

var err error

func main() {
	srv := server.New(localhost)
	ctx := context.TODO()

	storage.DB, err = storage.New(storageSqlitePath)
	if err != nil {
		log.Fatal(err)
	}

	if err = storage.DB.Init(ctx); err != nil {
		log.Fatal(err)
	}

	srv.Handler(handlers.UserGetHand())

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}

}
