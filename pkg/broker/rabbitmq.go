package broker

import (
	"fmt"
	"log"

	"github.com/Ajulll22/belajar-microservice/pkg/security"
	"github.com/streadway/amqp"
)

func RabbitMQConnect(RABBIT_HOST, RABBIT_USER, RABBIT_PASS, RABBIT_PORT string) RabbitMQ {
	clear_password := security.Decrypt(RABBIT_PASS, "62277ecdae08d9e813ab17a4ec2db8c58db38e398617824a2ef035c64d3da4be")

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", RABBIT_USER, clear_password, RABBIT_HOST, RABBIT_PORT)

	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return &rabbitMQ{
		Conn:    conn,
		Channel: channel,
	}
}

type RabbitMQ interface {
	DeclareExchange(exchangeName, exchangeType string) error
	DeclareQueue(queueName string, args amqp.Table) (amqp.Queue, error)
	BindQueue(queueName, exchangeName, routingKey string) error
	Publish(exchange, routingKey string, message []byte) error
	Consume(routes []ConsumerRoute) error
	Close()
}

type rabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (r *rabbitMQ) DeclareExchange(exchangeName, exchangeType string) error {
	return r.Channel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func (r *rabbitMQ) DeclareQueue(queueName string, args amqp.Table) (amqp.Queue, error) {
	return r.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		args,      // arguments
	)
}

func (r *rabbitMQ) BindQueue(queueName, exchangeName, routingKey string) error {
	return r.Channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
}

func (r *rabbitMQ) Publish(exchange, routingKey string, message []byte) error {
	err := r.Channel.Publish(
		exchange,   // Exchange
		routingKey, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	return err
}

func (r *rabbitMQ) Consume(routes []ConsumerRoute) error {
	for _, route := range routes {

		go func(route ConsumerRoute) {

			msgs, err := r.Channel.Consume(
				route.Queue,   // Queue name
				"",            // Consumer name
				route.AutoAck, // Auto-ack
				false,         // Exclusive
				false,         // No-local
				false,         // No-wait
				nil,           // Arguments
			)
			if err != nil {
				log.Println(route.Queue, "failed", err.Error())
				return
			}

			for msg := range msgs {

				if route.Async {
					go func(routeHandler ConsumerRoute, msg amqp.Delivery) {
						err := routeHandler.Handler(msg)
						if err != nil {
							log.Println("Error process message, ", err.Error())
						}
					}(route, msg)
				} else {
					err := route.Handler(msg)
					if err != nil {
						log.Println("Error process message, ", err.Error())
					}
				}
			}

		}(route)

	}

	return nil
}

func (r *rabbitMQ) Close() {
	r.Conn.Close()
	r.Channel.Close()
}

type ConsumerRoute struct {
	Key     string
	Handler func(amqp.Delivery) error
	Async   bool
	AutoAck bool
	Queue   string
}
