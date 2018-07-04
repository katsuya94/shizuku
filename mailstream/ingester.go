package main

import (
	"encoding/base64"

	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
)

const (
	User       = "me"
	ChaseQuery = "from:(no-reply@alertsp.chase.com) subject:(Your Single Transaction Alert from Chase)"
)

type MailstreamIngester struct {
	service *gmail.Service
}

func NewMailstreamIngester(service *gmail.Service) *MailstreamIngester {
	return &MailstreamIngester{service: service}
}

func (in *MailstreamIngester) Ingest(f func(message string) error) error {
	return in.service.Users.Messages.List(User).Q(ChaseQuery).Pages(context.Background(), func(res *gmail.ListMessagesResponse) error {
		for _, message := range res.Messages {
			body, err := in.fetchMessageBody(message.Id)
			if err != nil {
				return err
			}
			err = f(body)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (in *MailstreamIngester) fetchMessageBody(messageId string) (string, error) {
	message, err := in.service.Users.Messages.Get(User, messageId).Do()
	if err != nil {
		return "", err
	}
	b, err := base64.URLEncoding.DecodeString(message.Payload.Body.Data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
