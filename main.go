package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

const queueName = "shuffle.test"

func main() {

	user := flag.String("u", "guest", "Username for AMQP")
	pass := flag.String("p", "guest", "Password for AMQP")
	host := flag.String("host", "localhost", "Host for AMQP")
	port := flag.String("port", "5672", "Port for AMQP")
	message := flag.String("message", "test message", "custom message")

	flag.Parse()

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/", *user, *pass, *host, *port)
	log.Printf("Connection: %s", connString)

	conn, err := amqp.Dial(connString)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
		os.Exit(1)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %s", err)
		os.Exit(1)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %s", err)
		os.Exit(1)
	}

	for {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(*message),
			})
		log.Printf(" [x] Sent: %s", *message)
		if err != nil {
			log.Printf("Failed to publish a message: %v", err)
			os.Exit(1)
		}
	}
}
