package bootstrap

import (
	"time"

	"github.com/gojek/heimdall/v7/hystrix"
)

type HTTPOption struct {
	timeout          uint32
	numberOfRequests uint16
}

func NewClient(option *HTTPOption) *hystrix.Client {
	return hystrix.NewClient(
		hystrix.WithHTTPTimeout(time.Duration(option.timeout)),
		hystrix.WithHystrixTimeout(time.Duration(option.numberOfRequests)),
	)
}
