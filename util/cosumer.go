package util

import (
	"bytes"
	"encoding/json"
	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/luiz-otavio/ptero-go/io"
	"github.com/valyala/fastjson"
	"log"
	"net/http"
)

func Get(client *hystrix.Client, endpoint string, header http.Header) (*fastjson.Value, error) {
	response, err := client.Get(
		endpoint,
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return io.JSONBody(&response.Body)
}

func Post(client *hystrix.Client, endpoint string, header http.Header, body map[string]string) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Post(
		endpoint,
		bytes.NewReader(buf),
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func Patch(client *hystrix.Client, endpoint string, header http.Header, body map[string]string) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Patch(
		endpoint,
		bytes.NewReader(buf),
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func Delete(client *hystrix.Client, endpoint string, header http.Header) int {
	response, err := client.Delete(
		endpoint,
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func Put(client *hystrix.Client, endpoint string, header http.Header, body map[string]string) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Put(
		endpoint,
		bytes.NewReader(buf),
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}
