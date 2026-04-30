package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(brokers string) *Producer {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers),
	)
	if err != nil {
		log.Fatalf("failed to create kafka client: %v", err)
	}

	return &Producer{client: client}
}

func (p *Producer) Publish(topic string, msg interface{}) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	record := &kgo.Record{
		Topic: topic,
		Value: data,
	}

	p.client.Produce(context.Background(), record, func(_ *kgo.Record, err error) {
		if err != nil {
			log.Printf("failed to produce message: %v", err)
		}
	})

	return nil
}
