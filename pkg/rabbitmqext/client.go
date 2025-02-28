package rabbitmqext

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tanveerprottoy/event-processor-go/pkg/must"
)

var (
	instance *Client
	once     sync.Once
)

type Client struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func GetInstance(url string) *Client {
	once.Do(func() {
		instance = new(Client)
		instance.init(url)
	})

	return instance
}

func (c *Client) init(url string) {
	c.conn = must.Must(amqp.Dial(url))

	c.ch = must.Must(c.conn.Channel())
}

func (c *Client) Conn() *amqp.Connection {
	return c.conn
}

func (c *Client) Channel() *amqp.Channel {
	return c.ch
}

func (c *Client) Close() {
	c.ch.Close()

	c.conn.Close()
}
