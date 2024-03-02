package client

import (
	"github.com/gojek/heimdall/v7/hystrix"
	"github.com/luiz-otavio/pterodactyl-go-client/v2/api/client/account"
	"github.com/luiz-otavio/pterodactyl-go-client/v2/api/client/server"
	"github.com/luiz-otavio/pterodactyl-go-client/v2/api/client/system_permissions"
)

type PteroConnection interface {
	HttpClient() hystrix.Client
}

type PteroApplication interface {
	GetConnection() (PteroConnection, error)
}

type PteroClient interface {
	GetConnection() (PteroConnection, error)
	GetAccount() (account.ClientAccount, error)
	GetServers() ([]server.ClientServer, error)
	GetPermissions() (system_permissions.ClientPermissions, error)
}
