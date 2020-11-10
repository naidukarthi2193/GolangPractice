package main

import (
	"fmt"
	"strconv"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		err1 := ch.Publish(
			"",
			"TestQueue",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(strconv.Itoa(i)),
			},
		)
		if err1 != nil {
			panic(err1)
			break
		}
	}
	defer conn.Close()
	fmt.Println("Successfully Sent")
}
