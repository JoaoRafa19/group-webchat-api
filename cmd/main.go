package main

import (
	"sse/handlers"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	out := make(chan amqp.Delivery )
	// rabbitmqChannel, err := rabbitmq.OpenChannel()
	// if err != nil {
	// 	panic(err)
	// }

	// go rabbitmq.Consume("msgs", rabbitmqChannel, out) // só escreve a mensagem da fila na thread

	handlers.Serve(out)

}
