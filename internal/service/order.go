package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/ruanbekker/rbkr-ops-order-service/internal/model"
	"github.com/ruanbekker/rbkr-ops-order-service/internal/repository"
	"github.com/ruanbekker/rbkr-ops-order-service/internal/kafka"
	"github.com/ruanbekker/rbkr-ops-order-service/internal/metrics"
)

type OrderService struct {
	repo     *repository.OrderRepository
	producer *kafka.Producer
}

func NewOrderService(repo *repository.OrderRepository, producer *kafka.Producer) *OrderService {
	return &OrderService{
		repo:     repo,
		producer: producer,
	}
}

func (s *OrderService) CreateOrder(productID string, quantity int) (*model.Order, error) {
	order := &model.Order{
		ID:        uuid.New().String(),
		ProductID: productID,
		Quantity:  quantity,
		Status:    model.StatusCreated,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}

	metrics.OrdersCreated.Inc()

	_ = s.producer.Publish("order_created", map[string]interface{}{
		"order_id":   order.ID,
		"product_id": productID,
		"quantity":   quantity,
	})

	return order, nil
}

func (s *OrderService) GetOrders() ([]*model.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) GetOrder(id string) (*model.Order, error) {
	return s.repo.GetByID(id)
}

