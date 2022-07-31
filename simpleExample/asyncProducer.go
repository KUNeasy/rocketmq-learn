package simpleExample

import (
	"context"
	"log"
	"sync"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func SendAsyncMessage() {
	// 创建一个生产者实例
	producer, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),       //namesrc地址
		producer.WithGroupName("please_rename_unique_group_name"), // 生产者组名称
	)
	if err != nil {
		log.Fatalf("new producer err:%v\n", err)
	}
	producer.Start()

	// 相当于Java的CountDownLatch
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		// 创建消息实例,包括Topic, Message Body
		msg := &primitive.Message{
			Topic: "Jodie_topic_1023",
			Body:  []byte("Hello world"),
		}
		msg.WithTag("TagA")
		// 发送异步消息
		producer.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
			wg.Done()
			if err != nil {
				log.Printf("send async msg err : %v\n", err)
				return
			}
			log.Printf("%d OK %v\n", i, *result)
		}, msg)
	}

	// 阻塞直到wg的delta减到0
	wg.Wait()
	// 当不使用时,关闭producer实例
	producer.Shutdown()
}
