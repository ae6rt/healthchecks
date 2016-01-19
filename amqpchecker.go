package healthchecks

import (
	"fmt"
	"math/rand"

	"github.com/streadway/amqp"
)

type AMQPChecker struct {
	conn *amqp.Connection
}

func NewAMQPHealth(c *amqp.Connection) HealthChecker {
	return AMQPChecker{c}
}

func (checker AMQPChecker) Check() (bool, error) {
	ch, err := checker.conn.Channel()
	if err != nil {
		return false, err
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return false, err
	}

	defer func() {
		ch.QueueDelete(q.Name, false, false, true)
		ch.Close()
	}()

	body := fmt.Sprintf("%d", rand.Intn(100))
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		return false, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		true,   // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for d := range msgs {
		if string(d.Body) != body {
			return false, fmt.Errorf("AMQP health message expected %s but got %s\n", body, string(d.Body))
		}
		break
	}

	return true, nil
}
