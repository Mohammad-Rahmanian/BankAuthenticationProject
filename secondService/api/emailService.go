package api

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"time"
)

var (
	mailApiKey = "2d1c6793bd829e1db0c93e4aba3e5e2d-324e0bb2-46cbb1a3"
	domain     = "sandbox9037afd2c4234092bfe7684151bfae23.mailgun.org"
	sender     = "sajadrahmanian1616@gmail.com"
)

func SendMail(recipient string, state string) (string, error) {
	mg := mailgun.NewMailgun(domain, mailApiKey)
	body := fmt.Sprintf("Your authentication request status is %s", state)
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
