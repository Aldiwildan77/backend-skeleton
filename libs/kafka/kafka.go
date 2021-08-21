package kafka

import (
	"encoding/json"

	"github.com/Aldiwildan77/backend-skeleton/config"
	"github.com/Aldiwildan77/backend-skeleton/libs/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(topic, key string, payload interface{}) error {
	pByte, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var kByte []byte
	if key == "" {
		kByte = nil
	} else {
		kByte = []byte(key)
	}

	kfMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            kByte,
		Value:          pByte,
	}

	if err := config.KafkaProducer.Produce(kfMsg, nil); err != nil {
		logger.LogError(logger.ErrorLog{Error: err, Msg: "[Kafka Library][Produce] error while publish"})
		return err
	}

	return nil
}
