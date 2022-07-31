package simpleExample

import (
	"context"
	"log"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func SendSyncMessage() {
	// 创建一个生产者实例
	producer, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),       // namesrv 地址
		producer.WithGroupName("please_rename_unique_group_name"), // 生产者组名
	)
	if err != nil {
		log.Fatalf("new producer err : %v\n", err)
	}
	// 启动这个实例
	producer.Start()
	for i := 0; i < 100; i++ {
		// 创建一个消息实例,定一个Topic, message body
		msg := &primitive.Message{
			Topic: "TopicTest",
			Body:  []byte("Hello RocketMQ" + strconv.Itoa(i)),
		}
		// 添加tag
		msg.WithTag("TagA")
		// 调用SendSync方法把消息发送到其中一个broker中
		result, err := producer.SendSync(context.Background(), msg)
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}
		log.Printf("%v\n", *result)
	}
	// 当不在使用时,关闭生产者实例
	producer.Shutdown()
}
