package mailer

import (
	"github.com/scorredoira/email"
	"net/smtp"
	"os"
)

func SendTwoFactorMail(code string, to string) error {
	from := os.Getenv("MAILING_FROM")
	username := os.Getenv("MAILING_USERNAME")
	password := os.Getenv("MAILING_PASSWORD")
	host := os.Getenv("MAILING_HOST")
	port := os.Getenv("MAILING_PORT")

	message := email.NewHTMLMessage("2FA code", code)
	message.To = []string{to}
	message.From.Address = from
	message.BodyContentType = "text/html"

	var auth smtp.Auth
	if username != "" && password != "" {
		auth = smtp.PlainAuth("", username, password, host)
	}

	return email.Send(host+":"+port, auth, message)
}
