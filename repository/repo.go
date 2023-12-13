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
		"INSERT INTO delivery_order (name) VALUES ($1)",
		deliveryOrder.Name,
		//TODO: всё остальное
	)
	if err != nil {
		return 0, fmt.Errorf("exec: %w", err)
	}

	id, _ := res.LastInsertId()
	return id, nil
}
