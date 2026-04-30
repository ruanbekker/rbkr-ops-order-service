# rbkr-ops-order-service

Order Service for Order Processing System (OPS)

## About

OPS is a hobby microservice architecture project for my own experiments in Go.

## Structure

The home repository that acts as the documentation for this project is available at:

- [rbkr-order-processing-system-microservices](https://github.com/ruanbekker/rbkr-order-processing-system-microservices)

Directory structure:

```bash
├── cmd
│   └── api
│       └── main.go
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── db
│   │   └── postgres
│   │       └── postgres.go
│   ├── handler
│   │   └── order.go
│   ├── kafka
│   │   └── producer.go
│   ├── model
│   │   └── order.go
│   ├── repository
│   │   └── order.go
│   └── service
│       └── order.go
```

## Order Service

- REST API
- Postgres integration
- Kafka producer

## Build

```bash
go mod tidy
go build -o app ./cmd/api
```

<details>
  <summary>Click to expand</summary>

## Migrations

Install `golang-migrate` with homebrew.

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
mkdir -p migrations
migrate create -ext sql -dir migrations -seq create_orders_table
vim migrations/000001_create_orders_table.up.sql
vim migrations/000001_create_orders_table.down.sql
migrate -path migrations -database "postgres://user:pass@localhost:5432/orders?sslmode=disable" up
```

## Test

Create topic:

```bash
docker exec -it kafka kafka-topics.sh --create --topic order_created --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

Create order:

```bash
curl -X POST localhost:8080/orders -H "Content-Type: application/json" -d '{"product_id":"abc","quantity":2}'
```

Consume message from Kafka

```bash
docker exec -it kafka kafka-console-consumer.sh  --bootstrap-server localhost:9092 --topic order_created  --from-beginning

{"order_id":"bab0bf5c-7c3f-4170-9ed1-389c16e51bfe","product_id":"abc","quantity":2}
```

</details>
