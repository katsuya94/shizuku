package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/katsuya94/shizuku/common"
	gmail "google.golang.org/api/gmail/v1"
)

const (
	User       = "me"
	ChaseQuery = "from:(no-reply@alertsp.chase.com) subject:(Your Single Transaction Alert from Chase)"
)

var (
	NoMessagesError = fmt.Errorf("no more messages")
	ChaseRegexp     = regexp.MustCompile(`A charge of \(\$USD\) (?P<amount>\d+\.\d\d) at (?P<descriptor>.+) has been authorized on (?P<time>\d{2}/\d{2}/\d{4} \d{1,2}:\d{2}:\d{2} (AM|PM) [A-Z]+).`)
	ChaseTimeLayout = "01/02/2006 3:04:05 PM MST"
)

type MessageBodyParseError string

func (body MessageBodyParseError) Error() string {
	return fmt.Sprintf("failed to parse message body: %v", body)
}

type MailstreamIngester struct {
	service *gmail.Service
}

func NewMailstreamIngester(service *gmail.Service) *MailstreamIngester {
	return &MailstreamIngester{service: service}
}

func (ingester *MailstreamIngester) Ingest(ctx context.Context, f func(*common.Transaction) error) error {
	return ingester.service.Users.Messages.List(User).Q(ChaseQuery).Pages(ctx, func(res *gmail.ListMessagesResponse) error {
		if len(res.Messages) == 0 {
			return NoMessagesError
		}
		for _, message := range res.Messages {
			if err := ingester.processPartialMessage(ctx, message, f); err != nil {
				return err
			}
		}
		return nil
	})
}

func (ingester *MailstreamIngester) processPartialMessage(ctx context.Context, message *gmail.Message, f func(*common.Transaction) error) error {
	body, err := ingester.fetchMessageBody(ctx, message.Id)
	if err != nil {
		return err
	}
	transaction := &common.Transaction{Id: message.Id}
	if err := parseMessageBody(transaction, body); err != nil {
		return err
	}
	return f(transaction)
}

func (ingester *MailstreamIngester) fetchMessageBody(ctx context.Context, messageId string) (string, error) {
	message, err := ingester.service.Users.Messages.Get(User, messageId).Context(ctx).Do()
	if err != nil {
		return "", err
	}
	b, err := base64.URLEncoding.DecodeString(message.Payload.Body.Data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func parseMessageBody(transaction *common.Transaction, body string) error {
	submatch := ChaseRegexp.FindStringSubmatch(body)
	if submatch == nil {
		return MessageBodyParseError(body)
	}
	transaction.Amount = submatch[1]
	transaction.Description = submatch[2]
	t, err := time.ParseInLocation(ChaseTimeLayout, submatch[3], time.UTC)
	if err != nil {
		return err
	}
	transaction.Time = t.Format(time.RFC3339)
	return nil
}
