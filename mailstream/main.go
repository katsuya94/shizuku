package main

import (
	"log"

	"github.com/katsuya94/shizuku/common"
)

func main() {
	in := NewMailstreamIngester(GetService())
	err := in.Ingest(func(transaction *common.Transaction) error {
		log.Printf("%+v", transaction)
		return nil
	})
	if err != nil && err != NoMessagesError {
		log.Fatalf("Encountered error while processing messages: %v", err)
	}
}
