package main

import (
	"context"
	"log"
	"time"

	"github.com/katsuya94/shizuku/common"
	"github.com/robfig/cron"
)

func main() {
	c := cron.NewWithLocation(time.UTC)
	c.AddFunc("0 0 * * * *", deliverNewTransactions)
	c.Run()
}

func deliverNewTransactions() {
	ctx := context.Background()
	in := NewMailstreamIngester(GetService())
	err := in.Ingest(ctx, func(transaction *common.Transaction) error {
		log.Printf("%+v", transaction)
		return nil
	})
	if err != nil && err != NoMessagesError {
		log.Fatalf("Encountered error while processing messages: %v", err)
	}
}
