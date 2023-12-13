package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"wb_task1/models"
	"wb_task1/repository"
)

type IRepository interface {
	AddDeliveryOrder(ctx context.Context, deliveryOrder models.DeliveryOrder) (int64, error)
}

type App struct {
	nc *nats.Conn

	repo IRepository
}

func NewApp(db *sql.DB, nc *nats.Conn) *App {
	return &App{
		repo: repository.NewRepository(db),
		nc:   nc,
	}
}

func (a *App) Run(ctx context.Context) error {
	sub, err := a.nc.SubscribeSync("your_subject")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := sub.NextMsg(0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received a message: %s\n", msg.Data)

		var order models.Order
		err = json.Unmarshal(msg.Data, &order)
		if err != nil {
			return fmt.Errorf("unmarshal: %w", err)
		}

		dbOrder := models.DbOrder{
			//TODO: проинициализировать все поля, которые не доставка, оплата и айтемы
		}

		doId, err := a.repo.AddDeliveryOrder(ctx, order.DelOrder)
		if err != nil {
			return fmt.Errorf("add delivery order: %w", err)
		}
		dbOrder.DelOrderID = doId

		fmt.Printf("del order id = %d\n", doId)
		//TODO

		break // TODO DELETE

	}
	return nil
}
