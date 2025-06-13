package resend

import (
	"log"
	"os"

	"github.com/resend/resend-go/v2"
)

var client *resend.Client

func Init() {
	apiKey := os.Getenv("RESEND_KEY")
	if apiKey == "" {
		log.Fatal("RESEND_KEY não definida nas variáveis de ambiente")
	}

	client = resend.NewClient(apiKey)
}

// Send envia um e-mail usando o Resend
func Send(to []string, subject, html string) error {
	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      to,
		Subject: subject,
		Html:    html,
	}

	email, err := client.Emails.Send(params)
	if err != nil {
		log.Println("Erro ao enviar email:", err)
		return err
	}

	log.Println("Email enviado com sucesso! ID:", email.Id)
	return nil
}
