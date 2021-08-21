package registry

type KafkaConsumerProto interface {
	Start()
}

var _kafkaConsumers []KafkaConsumerProto

func RegisterKafkaConsumer(k KafkaConsumerProto) {
	_kafkaConsumers = append(_kafkaConsumers, k)
}

func LoadKafkaConsumers() []KafkaConsumerProto {
	return _kafkaConsumers
}
