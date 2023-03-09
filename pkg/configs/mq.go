package configs

import (
	"context"
	"mini-douyin/pkg/consts"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var MQP rocketmq.Producer

func initMQ() {
	endPoint := []string{consts.MQEndPoint}
	// 创建一个producer实例
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(endPoint),
		producer.WithRetry(2),
		producer.WithGroupName("mini-douyin"),
	)
	if err != nil {
		Logger.Error("RocketMQ Connect Fail" + err.Error())
		panic(err)
	}
	// 启动
	err = p.Start()
	if err != nil {
		Logger.Error("RocketMQ Connect Fail" + err.Error())
		panic(err)
	}
	MQP = p

	Logger.Info("RocketMQ Connect Success.")
}

func SendSyncMessage(topic string, msg string) error {
	_, err := MQP.SendSync(context.Background(), &primitive.Message{
		Topic: topic,
		Body:  []byte(msg),
	})
	if err != nil {
		Logger.Error(err.Error())
		return err
	}
	return nil
}
