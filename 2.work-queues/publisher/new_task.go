package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

type Biodata struct {
	Name string
	Age  int
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	log.Println("connecting to RabbitMQ")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Println("connecting to RabbitMQ channel")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	log.Println("declaring Queue")
	// q, err := ch.QueueDeclare(
	// 	"hello",
	// 	false, //durability
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )

	q2, err := ch.QueueDeclare(
		"durable_message",
		true, //durability
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare Queue")

	// membuat receiver tidak menerima message lebih dari 1
	// (sebelum receiver mengirimkan 1 acknowledge, jangan kirim pesan
	// ke receiver ini.)
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	log.Println("Publishing a  message")
	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",
		q2.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)
	failOnError(err, "failed to publish message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello....."
	} else {
		s = strings.Join(args[1:], " ")
	}

	return s
}
