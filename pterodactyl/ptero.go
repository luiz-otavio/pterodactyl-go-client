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

const (
	START = "start"
	STOP = "stop"
	RESTART = "restart"
	KILL = "kill"
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

func (client *PteroClient) Execute(uniqueId string, command string) int {
	body, err := json.Marshal(Body{
		"command": command,
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Post(
		client.endpoint("servers/"+uniqueId+"/command"),
		bytes.NewReader(body),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func (client *PteroClient) Rename(uniqueId string, name string) int {
	body, err := json.Marshal(Body{
		"name": name,
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Post(
		client.endpoint("servers/"+uniqueId+"/settings/rename"),
		bytes.NewReader(body),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func (client *PteroClient) Reinstall(uniqueId string) int {
	response, err := client.Connection.Post(
		client.endpoint("servers/"+uniqueId+"/settings/reinstall"),
		nil,
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func (client *PteroClient) Power(uniqueId string, power string) int {
	body, err := json.Marshal(Body{
		"signal": power,
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Connection.Post(
		client.endpoint("servers/"+uniqueId+"/power"),
		bytes.NewReader(body),
		client.header(),
	)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
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
