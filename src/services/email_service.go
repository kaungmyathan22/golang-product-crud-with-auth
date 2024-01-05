package services

import (
	"log"

	"gopkg.in/mail.v2"
)

type EmailService struct {
	smtpServer string
	smtpPort   int
	username   string
	password   string
	from       string
}

func NewEmailService() *EmailService {
	return &EmailService{
		smtpServer: "localhost",
		smtpPort:   1025,
		username:   "user@example.com",
		password:   "topsecret",
		from:       "Your App <yourapp@example.com>",
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := mail.NewDialer(e.smtpServer, e.smtpPort, e.username, e.password)

	err := d.DialAndSend(m)
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
