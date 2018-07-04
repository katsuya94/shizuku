package main

import (
	"log"
)

func main() {
	in := NewMailstreamIngester(getService())
	err := in.Ingest(func(body string) error {
		log.Print(body)
		return nil
	})
	if err != nil && err != NoMessagesError {
		log.Fatalf("Encountered error while processing messages: %v", err)
	}
}
