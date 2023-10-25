package api

import (
	"BankAuthenticationProject/configs"
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"time"
)

func SendMail(recipient string, state string) (string, error) {
	mg := mailgun.NewMailgun(configs.MailDomain, configs.MailApiKey)
	body := fmt.Sprintf("Your authentication request status is %s", state)
	m := mg.NewMessage(
		configs.MailSender,
		"Authentication Result",
		body,
		recipient,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}
