package kafka_consumer

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/laidingqing/amadd9/common/config"
)

func init() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               config.KafkaConsumer.Broker,
		"group.id":                        config.KafkaConsumer.Group,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": config.KafkaConsumer.OffsetReset}})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = c.SubscribeTopics(config.KafkaConsumer.Topics, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	StartConsumer(c, config.KafkaConsumer.ConsumerTimestampField)

}
