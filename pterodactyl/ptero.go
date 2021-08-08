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

type PteroClient struct {
	Key string
	URL string

	Client hystrix.Client
}

func NewConnection(url string, key string, option bootstrap.HTTPOption) *PteroClient {
	return &PteroClient{
		Key:    key,
		URL:    url,
		Client: *bootstrap.NewClient(option),
	}
}

func (client *PteroClient) Servers() (*fastjson.Value, error) {
	response, err := client.Client.Get(client.URL+"/api/application/servers", client.header())

	if err != nil {
		log.Fatalf(err.Error())
	}

	return io.JSONBody(&response.Body)
}

func (client *PteroClient) ServerById(id int) (*fastjson.Value, error) {
	response, err := client.Client.Get(client.URL+"api/application/servers/"+strconv.Itoa(id), client.header())

	if err != nil {
		log.Fatalf(err.Error())
	}

	return io.JSONBody(&response.Body)
}

func (client *PteroClient) ExternalServerById(id int) (*fastjson.Value, error) {
	response, err := client.Client.Get(client.URL+"api/application/servers/external/"+strconv.Itoa(id), client.header())

	if err != nil {
		log.Fatalf(err.Error())
	}

	return io.JSONBody(&response.Body)
}

func (client *PteroClient) ServerUpdate(id int, body map[string]string) int {
	buf, err := json.Marshal(body)

	if err != nil {
		log.Fatalf(err.Error())
	}

	response, err := client.Client.Patch(client.URL+"api/applications/servers/"+strconv.Itoa(id)+"/details", bytes.NewReader(buf), client.header())

	if err != nil {
		log.Fatalf(err.Error())
	}

	return response.StatusCode
}

func (client *PteroClient) header() http.Header {
	headers := make(http.Header)

	headers.Add("Accept", "application/json")
	headers.Add("Authorization", "Bearer "+client.Key)
	headers.Add("Content-Type", "application/json")

	return headers
}
