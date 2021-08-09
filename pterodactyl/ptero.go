package pterodactyl

import (
	http2 "github.com/luiz-otavio/ptero-go/http"
	"net/http"
	"strconv"

	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/luiz-otavio/ptero-go/bootstrap"
	"github.com/valyala/fastjson"
)

type PteroType int

const (
	Client PteroType = iota
	Application
)

const (
	START   = "start"
	STOP    = "stop"
	RESTART = "restart"
	KILL    = "kill"
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
	return http2.Get(
		client,
		client.endpoint("servers"),
		client.header(),
	)
}

func (client *PteroClient) ServerById(id int) (*fastjson.Value, error) {
	return http2.Get(
		client,
		client.endpoint("servers/"+strconv.Itoa(id)),
		client.header(),
	)
}

func (client *PteroClient) Execute(uniqueId string, command string) int {
	return http2.Post(
		client,
		client.endpoint("servers/"+uniqueId+"/command"),
		client.header(),
		Body{
			"command": command,
		})
}

func (client *PteroClient) Rename(uniqueId string, name string) int {
	return http2.Post(
		client,
		client.endpoint("servers/"+uniqueId+"/settings/rename"),
		client.header(),
		Body{
			"name": name,
		})
}

func (client *PteroClient) Reinstall(uniqueId string) int {
	return http2.Post(
		client,
		client.endpoint("servers/"+uniqueId+"/settings/reinstall"),
		client.header(),
		nil,
	)
}

func (client *PteroClient) Power(uniqueId string, power string) int {
	return http2.Post(
		client,
		client.endpoint("servers/"+uniqueId+"/power"),
		client.header(),
		Body{
			"signal": power,
		})
}

func (client *PteroClient) ExternalServerById(id int) (*fastjson.Value, error) {
	return http2.Get(
		client,
		client.endpoint("servers/external/"+strconv.Itoa(id)),
		client.header(),
	)
}

func (client *PteroClient) UpdateDetails(id int, body Body) int {
	return http2.Patch(
		client,
		client.endpoint("servers/"+strconv.Itoa(id)+"/details"),
		client.header(),
		body,
	)
}

func (client *PteroClient) UpdateInfo(id int, body Body) int {
	return http2.Patch(
		client,
		client.endpoint("servers/"+strconv.Itoa(id)+"/build"),
		client.header(),
		body,
	)
}

func (client *PteroClient) UpdateEnvironment(name string, body Body) int {
	return http2.Put(
		client,
		client.endpoint("servers/"+name+"/startup/variable"),
		client.header(),
		body,
	)
}

func (client *PteroClient) Create(info Body) int {
	return http2.Post(
		client,
		client.endpoint("servers"),
		client.header(),
		info,
	)
}

func (client *PteroClient) Delete(id int) int {
	return http2.Delete(
		client,
		client.endpoint("servers/"+strconv.Itoa(id)),
		client.header(),
	)
}

func (client *PteroClient) Environments(name string) (*fastjson.Value, error) {
	return http2.Get(
		client,
		client.endpoint("servers/"+name+"/startup"),
		client.header(),
	)
}

func (client *PteroClient) Resources(name string) (*fastjson.Value, error) {
	return http2.Get(
		client,
		client.endpoint("servers/"+name+"/resources"),
		client.header(),
	)
}

func (client *PteroClient) UpdateStartup(id int, body Body) int {
	return http2.Patch(
		client,
		client.endpoint("servers/"+strconv.Itoa(id)+"/startup"),
		client.header(),
		body,
	)
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
