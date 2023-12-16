package models

import (
	"github.com/lib/pq"
	"time"
)

type DeliveryOrder struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     int64  `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payments struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type DbOrder struct {
	DelOrderID        int64         `db:"delivery_id"`
	PaymentsID        int64         `db:"payment_id"`
	ItemsIDs          pq.Int64Array `db:"items"`
	OrderUid          string        `db:"uid"`
	TrackNumber       string        `db:"track_number"`
	Entry             string        `db:"entry"`
	Locale            string        `db:"locale"`
	InternalSignature string        `db:"internal_signature"`
	CustomerID        string        `db:"customer_id"`
	DeliveryService   string        `db:"delivery_service"`
	Shardkey          string        `db:"shardkey"`
	SmID              int           `db:"sm_id"`
	DateCreated       time.Time     `db:"date_created"`
	OofShard          string        `db:"oof_shard"`
}

type Order struct {
	OrderUid          string        `json:"order_uid"`
	TrackNumber       string        `json:"track_number"`
	Entry             string        `json:"entry"`
	DelOrder          DeliveryOrder `json:"delivery"`
	Payments          Payments      `json:"payment"`
	Items             []Item        `json:"items"`
	Locale            string        `json:"locale"`
	InternalSignature string        `json:"internal_signature"`
	CustomerID        string        `json:"customer_id"`
	DeliveryService   string        `json:"delivery_service"`
	Shardkey          string        `json:"shardkey"`
	SmID              int           `json:"sm_id"`
	DateCreated       time.Time     `json:"date_created"`
	OofShard          string        `json:"oof_shard"`
}
