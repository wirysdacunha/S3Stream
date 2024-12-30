package services

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/streadway/amqp"
)

type UploadMessage struct {
	FileName string `json:"fileName"`
	FileData string `json:"fileData"` // Base64 encoded file
}

func StartConsumer() error {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	log.Println(rabbitURL)
	if rabbitURL == "" {
		return fmt.Errorf("RABBITMQ_URL environment variable not set")
	}
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"S3Stream.jobs",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer: %v", err)
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("message received: %v", string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-sigchan

	log.Printf("interrupted, shutting down")
	forever <- true


	return nil
}