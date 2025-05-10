package mq

import (
	"diktok/config"
	"sync"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var p rocketmq.Producer
var c rocketmq.PushConsumer
var once sync.Once

func InitProducer(producerGroupName string) {
	var nameSrv, err = primitive.NewNamesrvAddr(config.System.MQ.Host + ":" + config.System.MQ.Port)
	if err != nil {
		panic(err.Error())
	}
	p, err = rocketmq.NewProducer(
		producer.WithNameServer(nameSrv),
		producer.WithGroupName(producerGroupName),
	)
	if err != nil {
		panic(err.Error())
	}
	p.Start()
}

func InitConsumer(consumerGroupName string) {
	var nameSrv, err = primitive.NewNamesrvAddr(config.System.MQ.Host + ":" + config.System.MQ.Port)
	if err != nil {
		panic(err.Error())
	}
	c, err = rocketmq.NewPushConsumer(
		consumer.WithGroupName(consumerGroupName),
		consumer.WithNameServer(nameSrv),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
	)
	if err != nil {
		panic(err.Error())
	}

	// c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	// 	for _, msg := range msgs {
	// 		fmt.Println(fmt.Sprintf("[消费消息] %s , %s", msg.MsgId, string(msg.Body)))
	// 	}
	// 	return consumer.ConsumeSuccess, nil
	// })

	// c.Start()
}

func GetProducer() rocketmq.Producer {
	if p == nil {
		once.Do(func() {
			InitProducer("default_group")
		})
	}
	return p
}

func GetConsumer() rocketmq.PushConsumer {
	if c == nil {
		once.Do(func() {
			InitConsumer("default_group")
		})
	}
	return c
}
