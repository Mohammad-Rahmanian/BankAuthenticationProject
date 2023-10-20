package api

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"time"
)

var (
	mailApiKey = "f3e9e9519c437e310904d3529be547c3-3750a53b-50e996c5"
	domain     = "sandbox109761c354b94e358b133277fa5d0881.mailgun.org"
	sender     = "sajad.rahmanian5070@gmail.com"
)

func SendMail(recipient string, state string) (string, error) {
	mg := mailgun.NewMailgun(domain, mailApiKey)
	body := fmt.Sprintf("Your authentication request status is %s.", state)
	m := mg.NewMessage(
		sender,
		"Authentication Result",
		body,
		recipient,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}
