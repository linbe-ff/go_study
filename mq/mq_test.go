package mq

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSimplePublish(t *testing.T) {
	rb := NewSimpleRabbitMQ("mallproduct")
	rb.SimplePublish("这个就是个消息生产者！！！")
}

func TestSimpleConsume(t *testing.T) {
	rb := NewSimpleRabbitMQ("mallproduct")
	rb.SimpleConsume()
}

func TestWorkPublish(t *testing.T) {
	rb := NewSimpleRabbitMQ("mallproduct")

	for i := 0; i < 100; i++ {
		rb.SimplePublish("这个就是个消息生产者！！！第" + strconv.Itoa(i) + "个消费者")
	}

	fmt.Println("发送成功......")
}

func TestWorkConsume(t *testing.T) {
	rb := NewSimpleRabbitMQ("mallproduct")
	rb.SimpleConsume()
}

func TestPubPublish(t *testing.T) {
	rb := NewPubSubRabbitMQ("iris-mall")

	for i := 0; i < 100; i++ {
		rb.PublishPub("订阅模式生产的第" + strconv.Itoa(i) + "条数据")
		fmt.Println("订阅模式生产的第" + strconv.Itoa(i) + "条数据")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("发送成功......")
}

func TestPubConsume(t *testing.T) {
	rb := NewPubSubRabbitMQ("iris-mall")
	rb.RecieveSub()
}

func TestRootingPublish(t *testing.T) {
	rb := NewRoutingRabbitMQ("amq.direct", "first")
	rb1 := NewRoutingRabbitMQ("amq.direct", "second")

	for i := 0; i < 100; i++ {
		rb.PublishRouting("订阅模式" + rb.Key + "生产的第" + strconv.Itoa(i) + "条数据")
		rb1.PublishRouting("订阅模式" + rb1.Key + "生产的第" + strconv.Itoa(i) + "条数据")
	}

	fmt.Println("发送成功......")
}

func TestRootingConsume(t *testing.T) {
	rb := NewRoutingRabbitMQ("amq.direct", "first")
	rb.RecieveRouting()
}

func TestRootingConsume2(t *testing.T) {
	rb := NewRoutingRabbitMQ("amq.direct", "second")
	rb.RecieveRouting()
}
