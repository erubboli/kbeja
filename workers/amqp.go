package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"github.com/erubboli/koding/metrics"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

var msgs <-chan amqp.Delivery

func Rabbit() <-chan amqp.Delivery {

	if msgs == nil {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		rabbit, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer rabbit.Close()

		q, err := rabbit.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when usused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		err = rabbit.QueueBind(
			q.Name, // queue name
			"",     // routing key
			"metrics", // exchange
			false,
			nil)

		failOnError(err, "Failed to bind a queue")

		msgs, err = rabbit.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

	}

	return msgs
}

func parseMessage(d amqp.Delivery) (metrics.Metric, error) {
	var m metrics.Metric
	err := json.Unmarshal(d.Body, m)
	return m, err
}
