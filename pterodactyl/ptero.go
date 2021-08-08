package pterodactyl

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/luiz-otavio/ptero-go/bootstrap"
	"github.com/luiz-otavio/ptero-go/io"
	"github.com/valyala/fastjson"
)

type PteroType int

const (
	Client PteroType = iota
	Application
)

type PteroClient struct {
	Key string
	URL string

	Connection hystrix.Client
	Type       PteroType
}

type Body map[string]string

func NewConnection(url string, key string, option bootstrap.HTTPOption, pteroType PteroType) *PteroClient {
	return &PteroClient{
		Key:        key,
		URL:        url,
		Connection: *bootstrap.NewClient(option),
		Type:       pteroType,
	}
}

func (client *PteroClient) Servers() (*fastjson.Value, error) {
	response, err := client.Connection.Get(
		client.endpoint("servers"),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return io.JSONBody(&response.Body)
}

func (client *PteroClient) ServerById(id int) (*fastjson.Value, error) {
	response, err := client.Connection.Get(
		client.endpoint("servers/"+strconv.Itoa(id)),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return io.JSONBody(&response.Body)
}

func (client *PteroClient) ExternalServerById(id int) (*fastjson.Value, error) {
	response, err := client.Connection.Get(
		client.endpoint("servers/external"+strconv.Itoa(id)),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return io.JSONBody(&response.Body)
}

func (client *PteroClient) UpdateDetails(id int, body Body) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Patch(
		client.endpoint("servers/"+strconv.Itoa(id)+"/details"),
		bytes.NewReader(buf),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func (client *PteroClient) UpdateInfo(id int, body Body) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Patch(
		client.endpoint("servers/"+strconv.Itoa(id)+"/build"),
		bytes.NewReader(buf),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func (client *PteroClient) UpdateStartup(id int, body Body) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Patch(
		client.endpoint("servers/"+strconv.Itoa(id)+"/startup"),
		bytes.NewReader(buf),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func (client *PteroClient) endpoint(target string) string {
	if client.Type == Client {
		return client.URL + "api/applications/" + target
	} else {
		return client.URL + "api/client/" + target
	}
}

func (client *PteroClient) header() http.Header {
	headers := make(http.Header)

	headers.Add("Accept", "application/json")
	headers.Add("Authorization", "Bearer "+client.Key)
	headers.Add("Content-Type", "application/json")

	return headers
}
