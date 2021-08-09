package http

import (
	"bytes"
	"encoding/json"
	"github.com/luiz-otavio/ptero-go/io"
	"github.com/luiz-otavio/ptero-go/pterodactyl"
	"github.com/valyala/fastjson"
	"log"
	"net/http"
)

func Get(client *pterodactyl.PteroClient, endpoint string, header http.Header) (*fastjson.Value, error) {
	response, err := client.Connection.Get(
		endpoint,
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return io.JSONBody(&response.Body)
}

func Post(client *pterodactyl.PteroClient, endpoint string, header http.Header, body pterodactyl.Body) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Post(
		endpoint,
		bytes.NewReader(buf),
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func Patch(client *pterodactyl.PteroClient, endpoint string, header http.Header, body pterodactyl.Body) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Patch(
		endpoint,
		bytes.NewReader(buf),
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func Put(client *pterodactyl.PteroClient, endpoint string, header http.Header, body pterodactyl.Body) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Put(
		endpoint,
		bytes.NewReader(buf),
		header,
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}
