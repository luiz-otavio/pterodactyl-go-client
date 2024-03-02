package server

import (
	"github.com/google/uuid"
	"github.com/luiz-otavio/pterodactyl-go-client/v2/api/client/schedules"
)

type PowerAction string

var (
	Start   PowerAction = "start"
	Stop    PowerAction = "stop"
	Restart PowerAction = "restart"
	Kill    PowerAction = "kill"
)

// TO-DO: Support websocket in the future
// TO-DO: Add support for file manager

type ClientServer interface {
	IsServerOwner() (bool, error)
	Identifier() (string, error)
	Uuid() (uuid.UUID, error)
	Name() (string, error)
	Node() (string, error)
	SFTP() (SFTP, error)
	Description() (string, error)
	Limits() (ServerLimit, error)
	FeatureLimits() (HostingLimit, error)
	Suspended() (bool, error)
	Installing() (bool, error)
	Relationships() ([]ClientAllocation, error)
	Usage() (ServerUsage, error)
	SendCommand(string) (bool, error)
	SendPower(PowerAction) (bool, error)
	Databases() ([]ServerDatabase, error)
	CreateDatabase(string, string) (ServerDatabase, error)
	DeleteDatabase(string) (bool, error)
	Schedules() ([]schedules.Schedule, error)
	Allocations() ([]ClientAllocation, error)
	AssignAllocation() (bool, error)
	SetAllocationNote(string, int) (bool, error)
	SetPrimaryAllocation(int) (bool, error)
	DeleteAllocation(int) (bool, error)
	Meta() (Meta, error)
}

type Meta struct {
	Owner       bool     `json:"is_server_owner"`
	Permissions []string `json:"user_permissions"`
}

type ClientAllocation interface {
	Id() (int, error)
	Ip() (string, error)
	Alias() (string, error)
	Port() (int, error)
	Notes() (string, error)
	Default() (bool, error)
}

type ServerUsage struct {
	CurrentState string `json:"current_state"`
	Suspended    bool   `json:"is_suspended"`
	Resources    struct {
		Memory  int `json:"memory_bytes"`
		Cpu     int `json:"cpu_absolute"`
		Disk    int `json:"disk_bytes"`
		RxBytes int `json:"network_rx_bytes"`
		TxBytes int `json:"network_tx_bytes"`
	}
}

type HostingLimit struct {
	Databases   int `json:"databases"`
	Allocations int `json:"allocations"`
	Backups     int `json:"backups"`
}

type ServerLimit struct {
	Memory int `json:"memory"`
	Swap   int `json:"swap"`
	Disk   int `json:"disk"`
	Io     int `json:"io"`
	Cpu    int `json:"cpu"`
}

type ServerDatabase struct {
	Id   string `json:"id"`
	Host struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	}
	Name           string `json:"name"`
	Username       string `json:"username"`
	Connections    string `json:"connections_from"`
	MaxConnections int    `json:"max_connections"`
}

type SFTP struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}
