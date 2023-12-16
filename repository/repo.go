package repository

import (
	"context"
	"database/sql"
	"fmt"

	"wb_task1/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddDeliveryOrder(ctx context.Context, deliveryOrder models.DeliveryOrder) (int64, error) {
	var id int64
	row := r.db.QueryRow(
		"INSERT INTO delivery_orders (name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		deliveryOrder.Name,
		deliveryOrder.Phone,
		deliveryOrder.Zip,
		deliveryOrder.City,
		deliveryOrder.Address,
		deliveryOrder.Region,
		deliveryOrder.Email,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("exec: %w", err)
	}

	return id, nil
}

func (r *Repository) AddPayments(ctx context.Context, payments models.Payments) (int64, error) {
	var id int64
	row := r.db.QueryRow(
		"INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		payments.Transaction,
		payments.RequestID,
		payments.Currency,
		payments.Provider,
		payments.Amount,
		payments.PaymentDt,
		payments.Bank,
		payments.DeliveryCost,
		payments.GoodsTotal,
		payments.CustomFee,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("exec: %w", err)
	}

	return id, nil
}

func (r *Repository) AddItems(ctx context.Context, items []models.Item) ([]int64, error) {
	var ids []int64
	for _, item := range items {
		var id int64
		row := r.db.QueryRow(
			"INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status,
		)
		err := row.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("exec: %w", err)
		}

		ids = append(ids, id)
	}
	return ids, nil
}

func (r *Repository) AddOrder(ctx context.Context, order models.DbOrder) (int64, error) {
	var id int64
	row := r.db.QueryRow(
		"INSERT INTO orders (uid, track_number, entry, delivery_id, payment_id, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id",
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		order.DelOrderID,
		order.PaymentsID,
		order.ItemsIDs,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("exec: %w", err)
	}

	return id, nil
}
