package main

import (
	"fmt"

	"github.com/streadway/amqp" // Import the amqp package
)

func main() {

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Close the connection when the function returns
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume("Just a hello", "", true, false, false, false, nil) // Consume the queue

	forever := make(chan bool)

	// Start a goroutine to listen for messages
	go func() {
		for d := range msgs { // Loop through the messages
			fmt.Println(string(d.Body)) // Print the message
		}
	}()

	fmt.Println("Waiting for messages...") // Wait for messages
	<-forever

	fmt.Println("Done") // When the loop is done, print "Done"
}
