package main

import (
	"fmt"

	"github.com/streadway/amqp" // Import the amqp package
)

func main() {

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")

	if err != nil {
		panic(err)
	}

	// Close the connection when the function returns
	defer conn.Close()

	fmt.Println("Connected")

	// Open a channel
	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	// Close the channel when the function returns
	defer ch.Close()

	// Declare a queue
	qu, err := ch.QueueDeclare("Just a hello", false, false, false, false, nil)

	fmt.Println(qu)

	if err != nil {
		panic(err)
	}

	fmt.Println("Queue declared")

	// Publish a message to the queue
	err = ch.Publish("", qu.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Another"),
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Successfuly Published Message")

	fmt.Println("Message sent")

}
