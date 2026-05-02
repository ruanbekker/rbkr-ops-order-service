package main

import (
    "github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

    "github.com/ruanbekker/rbkr-ops-order-service/internal/db/postgres"
    "github.com/ruanbekker/rbkr-ops-order-service/internal/handler"
    "github.com/ruanbekker/rbkr-ops-order-service/internal/kafka"
    "github.com/ruanbekker/rbkr-ops-order-service/internal/metrics"
    "github.com/ruanbekker/rbkr-ops-order-service/internal/repository"
    "github.com/ruanbekker/rbkr-ops-order-service/internal/service"
    "github.com/ruanbekker/rbkr-ops-order-service/pkg/config"
)

func main() {
    cfg := config.Load()

    db := postgres.New(cfg.DatabaseURL)
    producer := kafka.NewProducer(cfg.KafkaBrokers)

    repo := repository.NewOrderRepository(db)
    svc := service.NewOrderService(repo, producer)
    h := handler.NewOrderHandler(svc)

    metrics.Init()

    r := gin.Default()
    r.POST("/orders", h.CreateOrder)
    r.GET("/orders", h.GetOrders)
    r.GET("/orders/:id", h.GetOrder)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/healthz", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    r.Run(":8080")
}
