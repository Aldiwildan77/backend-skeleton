package cons

import "errors"

var (
	ConsumerError = errors.New("an error occur at integrator harga modal consume")
	RetryLimit    = errors.New("limit exceeded")
)
