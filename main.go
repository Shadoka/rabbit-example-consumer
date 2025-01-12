package main

import (
	"encoding/json"
	"fmt"
	"os"
	"rabbit-example-consumer/data"
	"rabbit-example-consumer/rabbit"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn := rabbit.Connect("localhost", "5672", "guest", "guest")
	defer conn.Close()
	rabbitChan := rabbit.GetChannel(conn)
	defer rabbitChan.Close()

	workerChannel := make(chan string)

	go ConsumeWithoutReply(rabbitChan, "addition-no-reply", workerChannel)
	go ConsumeWithReply(rabbitChan, "addition", "reply", workerChannel)

	for workerMsg := range workerChannel {
		fmt.Println(workerMsg)
	}
}

func ConsumeWithoutReply(rabbitChannel *amqp.Channel, queueName string, workerChannel chan string) {
	msgs, err := rabbitChannel.Consume(queueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot create consumer: %v", err)
		return
	}

	for msg := range msgs {
		workerChannel <- fmt.Sprintf("received msg with routing key '%v'", msg.RoutingKey)
	}
}

func ConsumeWithReply(rabbitChannel *amqp.Channel, queueName string, outputQueueName string, workerChannel chan string) {
	msgs, err := rabbitChannel.Consume(queueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "cannot create consumer: %v", err)
		return
	}

	for msg := range msgs {
		var inputMsg data.AdditionInput
		err = json.Unmarshal(msg.Body, &inputMsg)
		if err != nil {
			workerChannel <- fmt.Sprintf("cannot parse input message: %v", err)
			continue
		}

		additionResult := data.AdditionResult{
			Result:    inputMsg.FirstNumber + inputMsg.SecondNumber,
			RequestId: inputMsg.RequestId,
		}
		serializedResult, err := json.Marshal(additionResult)
		if err != nil {
			workerChannel <- fmt.Sprintf("cannot serialize addition result: %v", err)
		}

		err = rabbitChannel.Publish("addition-reply", "#", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        serializedResult,
			MessageId:   msg.MessageId,
		})
		if err != nil {
			workerChannel <- fmt.Sprintf("error publishing message: %v", err)
		}

		workerChannel <- fmt.Sprintf("received msg with routing key '%v'", msg.RoutingKey)
	}
}