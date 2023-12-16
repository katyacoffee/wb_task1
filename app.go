package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/nats-io/nats.go"
	"wb_task1/models"
	"wb_task1/repository"
)

type IRepository interface {
	AddDeliveryOrder(ctx context.Context, deliveryOrder models.DeliveryOrder) (int64, error)
	AddPayments(ctx context.Context, payments models.Payments) (int64, error)
	AddItems(ctx context.Context, items []models.Item) ([]int64, error)
	AddOrder(ctx context.Context, order models.DbOrder) (int64, error)
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
	//sub, err := a.nc.SubscribeSync("your_subject")
	//if err != nil {
	//	log.Fatal(err)
	//}

	for {
		//msg, err := sub.NextMsg(0)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("Received a message: %s\n", msg.Data)
		//
		//var order models.Order
		//err = json.Unmarshal(msg.Data, &order)
		//if err != nil {
		//	return fmt.Errorf("unmarshal: %w", err)
		//}

		order := models.Order{
			OrderUid: "LOLKEK",
			DelOrder: models.DeliveryOrder{
				Name: "delivery lol",
			},
			Payments: models.Payments{
				Currency: "SHEKELI",
			},
			Items: []models.Item{
				{
					Price: 100500,
				},
				{
					Price: 12345,
				},
			},
		}

		dbOrder := models.DbOrder{
			OrderUid:          order.OrderUid,
			TrackNumber:       order.TrackNumber,
			Entry:             order.Entry,
			Locale:            order.Locale,
			InternalSignature: order.InternalSignature,
			CustomerID:        order.CustomerID,
			DeliveryService:   order.DeliveryService,
			Shardkey:          order.Shardkey,
			SmID:              order.SmID,
			DateCreated:       order.DateCreated,
			OofShard:          order.OofShard,
		}

		doId, err := a.repo.AddDeliveryOrder(ctx, order.DelOrder)
		if err != nil {
			return fmt.Errorf("add delivery order: %w", err)
		}
		dbOrder.DelOrderID = doId

		fmt.Printf("del order id = %d\n", doId)

		pId, err := a.repo.AddPayments(ctx, order.Payments)
		if err != nil {
			return fmt.Errorf("add delivery order: %w", err)
		}
		dbOrder.PaymentsID = pId

		fmt.Printf("payment id = %d\n", pId)

		itIds, err := a.repo.AddItems(ctx, order.Items)
		if err != nil {
			return fmt.Errorf("add delivery order: %w", err)
		}
		dbOrder.ItemsIDs = itIds

		fmt.Printf("items ids = %v\n", itIds)

		oID, err := a.repo.AddOrder(ctx, dbOrder)
		if err != nil {
			return fmt.Errorf("add order: %w", err)
		}

		fmt.Printf("order id = %d\n", oID)

		break // TODO DELETE

	}
	return nil
}
