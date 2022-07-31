package simpleExample

import (
	"context"
	"log"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func SendOneWayMessage() {
	producer, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithGroupName("please_rename_unique_group_name"),
	)
	if err != nil {
		log.Fatalf("new producer err:%v\n", err)
	}
	producer.Start()

	for i := 0; i < 100; i++ {
		msg := &primitive.Message{
			Topic: "TopicTest",
			Body:  []byte("Hello RocketMQ" + strconv.Itoa(i)),
		}
		err := producer.SendOneWay(context.Background(), msg)
		if err != nil {
			log.Printf("send one way err :%v\n", err)
			continue
		}
		log.Println("send one way successfully")
	}
	producer.Shutdown()
}
