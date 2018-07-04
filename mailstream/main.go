package main

import (
	"encoding/base64"
	"log"

	// "golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
)

const (
	chaseQuery = "from:(no-reply@alertsp.chase.com) subject:(Your Single Transaction Alert from Chase)"
)

func process(message *gmail.Message) error {
	log.Printf("%+v", message)
	b, err := base64.URLEncoding.DecodeString(message.Payload.Body.Data)
	if err != nil {
		return err
	}
	body := string(b)
	log.Print(body)
	return nil
}

func main() {
	user := "me"
	srv := getService()
	// err := srv.Users.Messages.List(user).Q(chaseQuery).Pages(context.Background(), func(res *gmail.ListMessagesResponse) error {
	// 	for _, message := range res.Messages {
	// 		process(message)
	// 	}
	// 	return nil
	// })
	res, err := srv.Users.Messages.List(user).Q(chaseQuery).Do()
	if err != nil {
		log.Fatalf("Encountered error while processing messages: %v", err)
	}
	for _, message := range res.Messages {
		process(message)
	}
}
