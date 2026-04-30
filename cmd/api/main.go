package main

import (
    "github.com/gin-gonic/gin"

    "github.com/ruanbekker/rbkr-order-processing-system-microservices/internal/db/postgres"
    "github.com/ruanbekker/rbkr-order-processing-system-microservices/internal/handler"
    "github.com/ruanbekker/rbkr-order-processing-system-microservices/internal/kafka"
    "github.com/ruanbekker/rbkr-order-processing-system-microservices/internal/repository"
    "github.com/ruanbekker/rbkr-order-processing-system-microservices/internal/service"
    "github.com/ruanbekker/rbkr-order-processing-system-microservices/pkg/config"
)

func main() {
    cfg := config.Load()

    db := postgres.New(cfg.DatabaseURL)
    producer := kafka.NewProducer(cfg.KafkaBrokers)

    repo := repository.NewOrderRepository(db)
    svc := service.NewOrderService(repo, producer)
    h := handler.NewOrderHandler(svc)

    r := gin.Default()
    r.POST("/orders", h.CreateOrder)
    r.GET("/orders", h.GetOrders)
    r.GET("/orders/:id", h.GetOrder)
	r.GET("/healthz", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    r.Run(":8080")
}
