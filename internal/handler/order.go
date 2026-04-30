package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruanbekker/rbkr-order-processing-system-microservices/internal/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{service: svc}
}

type CreateOrderRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	order, err := h.service.CreateOrder(req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	orders, err := h.service.GetOrders()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.service.GetOrder(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	c.JSON(200, order)
}
