package repository

import (
	"database/sql"

	"github.com/ruanbekker/rbkr-order-processing-system-microservices/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *model.Order) error {
	query := `
		INSERT INTO orders (id, product_id, quantity, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		query,
		order.ID,
		order.ProductID,
		order.Quantity,
		order.Status,
		order.CreatedAt,
	)

	return err
}

func (r *OrderRepository) GetAll() ([]*model.Order, error) {
	query := `
		SELECT id, product_id, quantity, status, created_at
		FROM orders
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*model.Order

	for rows.Next() {
		var o model.Order
		if err := rows.Scan(
			&o.ID,
			&o.ProductID,
			&o.Quantity,
			&o.Status,
			&o.CreatedAt,
		); err != nil {
			return nil, err
		}

		orders = append(orders, &o)
	}

	return orders, nil
}

func (r *OrderRepository) GetByID(id string) (*model.Order, error) {
	query := `
		SELECT id, product_id, quantity, status, created_at
		FROM orders
		WHERE id = $1
	`

	var o model.Order

	err := r.db.QueryRow(query, id).Scan(
		&o.ID,
		&o.ProductID,
		&o.Quantity,
		&o.Status,
		&o.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &o, nil
}

