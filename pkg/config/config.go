package config

import "os"

type Config struct {
	DatabaseURL  string
	KafkaBrokers string
}

func Load() *Config {
	return &Config{
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/orders?sslmode=disable"),
		KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

