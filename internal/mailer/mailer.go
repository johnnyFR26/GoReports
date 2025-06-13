package mailer

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(recipients []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "s4V8o@example.com")
	m.SetHeader("To", recipients...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "s4V8o@example.com", "password")

	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	return nil
}
