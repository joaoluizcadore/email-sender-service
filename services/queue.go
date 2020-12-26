package services

import (
	"log"

	"github.com/streadway/amqp"
)

//StartQueue - Initialize the queue and call the MessageHandlerType on every message received
func StartQueue(fnc MessageHandlerType) {
	config := GetConfig()
	conn, err := amqp.Dial(config.Queue.Host)
	ExitOnFail(err, "Cannot open connection with Queue broker")

	defer conn.Close()

	ch, err := conn.Channel()
	ExitOnFail(err, "Cannot open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.Queue.Name, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//log.Printf("Received a message: %s", d.Body)
			fnc(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages... ")
	<-forever
}
