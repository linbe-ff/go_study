package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// 生产者
func (r *RabbitMQ) PublishPub(message string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.GetError(err, "RabbitMQ创建exchange失败！")

	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.GetError(err, "消费者创建失败")
}

// 消费者
func (r *RabbitMQ) RecieveSub() {
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.GetError(err, "RabbitMQ创建exchange失败！")

	//创建队列
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.GetError(err, "创建队列失败")

	//交换机和队列绑定
	err = r.channel.QueueBind(
		//这里指定的是队列名称，由于上面设置的队列是空的代表随机，这里要获取名字，只能从q中取
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)

	//消费者
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	waitChannel := make(chan bool)
	go func() {
		for msg := range messages {
			fmt.Println(string(msg.Body))
		}
	}()
	<-waitChannel
}
