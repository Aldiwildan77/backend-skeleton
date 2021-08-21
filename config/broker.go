package config

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	// Kafka represent kafka config
	Kafka *kafka.ConfigMap

	// KafkaProducer represent kafka consumer
	KafkaConsumer *kafka.Consumer

	// KafkaProducer represent kafka producer
	KafkaProducer *kafka.Producer
)

func KafkaConsumerInstance() *kafka.Consumer {
	conf := loadKafkaConfiguration()
	srvs := ""

	if conf.MultipleHost != "" {
		srvs = conf.MultipleHost
	} else {
		srvs = fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	}

	kfConfig := &kafka.ConfigMap{
		"bootstrap.servers":    srvs,
		"group.id":             conf.GroupID,
		"enable.partition.eof": true,
	}

	c, err := kafka.NewConsumer(kfConfig)
	if err != nil {
		panic(err)
	}

	return c
}

func KafkaProducerInstance() *kafka.Producer {
	conf := loadKafkaConfiguration()
	srvs := ""

	if conf.MultipleHost != "" {
		srvs = conf.MultipleHost
	} else {
		srvs = fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	}

	kfConfig := &kafka.ConfigMap{
		"bootstrap.servers": srvs,
	}

	c, err := kafka.NewProducer(kfConfig)
	if err != nil {
		panic(err)
	}

	return c
}

func loadKafkaConfiguration() KafkaConfig {
	conf := KafkaConfig{
		Host:    Cfg.Kafka.Host,
		Port:    Cfg.Kafka.Port,
		GroupID: Cfg.Kafka.GroupID,
	}

	return conf
}
