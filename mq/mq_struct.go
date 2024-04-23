package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const (
	USER     = "guest"
	PASSWORD = "guest"
	HOST     = "127.0.0.1"
	PORT     = "5674"
)

type (
	RabbitMQ struct {
		conn      *amqp.Connection
		channel   *amqp.Channel
		QueueName string
		Exchange  string
		Key       string
		Path      string
	}
)

func (r *RabbitMQ) ToPath() {
	r.Path = fmt.Sprintf("amqp://%s:%s@%s:%s/", USER, PASSWORD, HOST, PORT)
}

// 实例化
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
	}
	rabbitmq.ToPath()
	var err error

	rabbitmq.conn, err = amqp.Dial(rabbitmq.Path)
	rabbitmq.GetError(err, "RabbitMQ连接不正确！")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.GetError(err, "RabbitMQ连接channel不正确！")
	return rabbitmq
}

// 订阅模式创建实例化
func NewPubSubRabbitMQ(exchangeName string) *RabbitMQ {
	var err error
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Path)
	rabbitmq.GetError(err, "创建conn失败")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.GetError(err, "创建channel失败")
	return rabbitmq
}

// 获取错误
func (r *RabbitMQ) GetError(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", err, message)
		panic(fmt.Sprintf("%s:%s", err, message))
	}
}
