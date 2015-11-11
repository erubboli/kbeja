package main

import (
	"fmt"
	"os"
  "log"
	"encoding/json"
  "math/rand"

	"github.com/streadway/amqp"
	"github.com/erubboli/kbeja/metrics"
)


func SendMessage() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"metrics",// name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	err = ch.Publish(
		"metrics", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        bodyFrom(os.Args),
		})

	failOnError(err, "Failed to publish a message")

	log.Printf("Message Sent.")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func bodyFrom(args []string) []byte{
  var m metrics.Metric
  if len(args) > 2 {
    m.Username = args[2]
  }
  if len(args) > 3 {
    m.Metric = args[3]
  }
  m.Count = rand.Int63()
  body, _ := json.Marshal(m)
  return body
}
