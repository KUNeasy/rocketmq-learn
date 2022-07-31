package simpleExample

import (
	"context"
	"log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func ConsumeMessage() {
	// 创建一个push客户端,消息队列主动发送消息给客户端,如果客户端可以接收消息,则发送消息给客户端,如果客户端不能处理消息,消息队列则继续请求,直到客户端可以接受消息
	// pull模式,客户端轮循请求消息队列看是否有消息,如果有消息则拉取,处理消息,处理完消息继续请求消息队列;没有消息则继续轮循
	csr, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("please_rename_unique_group_name"),
	)
	if err != nil {
		log.Fatalf("new consumer err :%v\n", err)
	}
	csr.Subscribe("TopicTest", consumer.MessageSelector{}, func(ctx context.Context, me ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range me {
			log.Printf("consume message :%v\n", msg)
		}
		return consumer.ConsumeSuccess, nil
	})
	csr.Start()
}
