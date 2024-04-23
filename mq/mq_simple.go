package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 销毁连接
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 简单模式
func NewSimpleRabbitMQ(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式生产者
func (r *RabbitMQ) SimplePublish(message string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // durable  是否持久化数据
		false, // autoDelete 是否自动删除
		false, // exclusive 排他性，权限私有
		false, // noWaite 是否阻塞
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, //mandatory 如果是true，会根据exchange和routekey规则，如果无法找到符合条件的队列会把消息的消息返回给发送者
		false, //immediate true 表示当exchange将消息发送到队列后发现队列没有绑定消费者，则会把消息发回给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

// 简单模式消费者
func (r *RabbitMQ) SimpleConsume() {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // durable  是否持久化数据
		false, // autoDelete 是否自动删除
		false, // exclusive 排他性，权限私有
		false, // noWaite 是否阻塞
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	message, err := r.channel.Consume(
		r.QueueName,
		"",    //区分多个消费者
		true,  //autoAck 是否自动应答
		false, //exclusive 是否排它性
		false, //noLocal true表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	channelWait := make(chan bool)
	go func() {
		for msg := range message {
			fmt.Println(string(msg.Body))
		}
	}()

	log.Printf("[*] waiting......")
	<-channelWait
}
