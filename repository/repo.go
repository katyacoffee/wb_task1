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
	res, err := r.db.Exec(
		"INSERT INTO delivery_orders (name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		deliveryOrder.Name,
		deliveryOrder.Phone,
		deliveryOrder.Zip,
		deliveryOrder.City,
		deliveryOrder.Address,
		deliveryOrder.Region,
		deliveryOrder.Email,
	)
	if err != nil {
		return 0, fmt.Errorf("exec: %w", err)
	}

	id, _ := res.LastInsertId() //TODO: разобраться, почему тут 0
	return id, nil
}
