package main

import (
	"context"
	"database/sql"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=12345 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Подключение к серверу NATS Streaming
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	app := NewApp(db, nc)
	err = app.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
