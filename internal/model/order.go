package model

import "time"

type OrderStatus string

const (
	StatusCreated OrderStatus = "created"
	StatusFailed  OrderStatus = "failed"
	StatusDone    OrderStatus = "done"
)

type Order struct {
	ID        string      `json:"id" db:"id"`
	ProductID string      `json:"product_id" db:"product_id"`
	Quantity  int         `json:"quantity" db:"quantity"`
	Status    OrderStatus `json:"status" db:"status"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
}
