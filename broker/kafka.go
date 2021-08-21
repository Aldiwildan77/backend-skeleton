package broker

import (
	"github.com/Aldiwildan77/backend-skeleton/app/registry"
)

func KafkaInstance() {
	// kafka.Register()

	consumers := registry.LoadKafkaConsumers()
	for _, c := range consumers {
		c.Start()
	}
}
