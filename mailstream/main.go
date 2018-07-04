package main

import (
	"log"

	"github.com/katsuya94/shizuku/common"
)

func main() {
	in := NewMailstreamIngester(getService())
	err := in.Ingest(func(t *common.Transaction) error {
		log.Printf("%+v", t)
		return nil
	})
	if err != nil && err != NoMessagesError {
		log.Fatalf("Encountered error while processing messages: %v", err)
	}
}
