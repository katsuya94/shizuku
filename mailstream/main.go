package main

import (
	"context"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/katsuya94/shizuku/common"
	"github.com/robfig/cron"
)

const (
	topic = "transactions"
)

type messageEncoder struct {
	message proto.Message
}

func (e messageEncoder) Encode() ([]byte, error) {
	return proto.Marshal(e.message)
}

func (e messageEncoder) Length() int {
	b, _ := e.Encode()
	return len(b)
}

func main() {
	c := cron.NewWithLocation(time.UTC)
	c.AddFunc("0 0 * * * *", deliverNewTransactions)
	c.Run()
}

func deliverNewTransactions() {
	ctx := context.Background()
	ingester := NewMailstreamIngester(GetService())
	producer, err := sarama.NewSyncProducer([]string{"kafka"}, nil)
	if err != nil {
		log.Fatalf("error starting producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("error closing producer: %v", err)
		}
	}()
	err = ingester.Ingest(ctx, func(transaction *common.Transaction) error {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: messageEncoder{transaction},
		}
		_, _, err := producer.SendMessage(message)
		return err
	})
	if err != nil && err != NoMessagesError {
		log.Fatalf("error ingesting messages: %v", err)
	}
}
