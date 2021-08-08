package bootstrap

import (
	"time"

	"github.com/gojek/heimdall/v7/hystrix"
)

type HTTPOption struct {
	Timeout     int64
	RequestSize uint16
}

func NewClient(option HTTPOption) *hystrix.Client {
	return hystrix.NewClient(
		hystrix.WithHTTPTimeout(time.Duration(option.Timeout)),
		hystrix.WithHystrixTimeout(time.Duration(option.RequestSize)),
	)
}
