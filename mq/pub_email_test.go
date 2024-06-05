package mq

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"net/smtp"
	"os"
	"text/template"
)

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

func sendEmail(email Email) error {
	t, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := t.Execute(&body, email); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))

	err = smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, email.From, []string{email.To}, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email_queue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			email := Email{
				From:    os.Getenv("SMTP_FROM"),
				To:      string(d.Body),
				Subject: "Test Email",
				Body:    "This is a test email sent from the email system.",
			}

			if err := sendEmail(email); err != nil {
				log.Printf("Failed to send email: %v", err)
			} else {
				log.Printf("Email sent to: %s", email.To)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
