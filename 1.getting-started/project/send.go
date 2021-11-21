package main

import (
	"log"

	"github.com/streadway/amqp"
)

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
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "failed to declare Queue")

	log.Println("Publishing a  message")
	body := "Hello, Fany Ervansyah"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	failOnError(err, "failed to publish message")

}
