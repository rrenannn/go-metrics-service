package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Publisher struct {
	conn     *AMQPConn
	exchange string
}

type AMQPConn struct {
	conn *amqp.Connection
}

type Config struct {
	ExchangeName string
}

func NewAMQPConn(url string) (*AMQPConn, error) { // <- aonde abre a conexao com o rabbitmq
	conn, err := amqp.DialConfig(url, amqp.Config{
		Dial: amqp.DefaultDial(30 * time.Second),
	})
	if err != nil {
		return nil, err
	}
	return &AMQPConn{conn: conn}, nil
}

func NewPublisher(conn *AMQPConn, cfg Config) *Publisher {
	return &Publisher{conn: conn, exchange: cfg.ExchangeName}
}

func (a *AMQPConn) Channel() (*amqp.Channel, error) {
	return a.conn.Channel() // <- canal onde ira receber as mensagens
}

func (p *Publisher) Publish(routingKey string, body []byte) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare(p.exchange, "direct", true, false, false, false, nil); err != nil {
		return err
	}

	return ch.Publish(p.exchange, routingKey, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         body,
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
	})
}
